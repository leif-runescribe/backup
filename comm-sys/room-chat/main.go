// main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type ChatMode int

const (
	ModeLobby ChatMode = iota
	ModePrivateChat
	ModeGroupChat
)

type Message struct {
	from    string
	content string
	time    time.Time
}

type Group struct {
	name     string
	members  map[string]*Client
	messages []Message
	mutex    sync.RWMutex
}

type Client struct {
	id       string
	name     string
	conn     net.Conn
	outbound chan string
	paired   *Client
	group    *Group
	mode     ChatMode
	server   *Server
}

type Server struct {
	clients  map[string]*Client
	groups   map[string]*Group
	mutex    sync.RWMutex
	requests map[string]string // map[fromID]toID
}

func NewServer() *Server {
	return &Server{
		clients:  make(map[string]*Client),
		groups:   make(map[string]*Group),
		requests: make(map[string]string),
	}
}

func (s *Server) createGroup(name string) *Group {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	group := &Group{
		name:    name,
		members: make(map[string]*Client),
	}
	s.groups[name] = group
	return group
}

func (g *Group) broadcast(msg Message, exclude *Client) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	g.messages = append(g.messages, msg)
	for _, member := range g.members {
		if member != exclude {
			member.outbound <- fmt.Sprintf("[%s] %s: %s",
				msg.time.Format("15:04:05"),
				msg.from,
				msg.content)
		}
	}
}

func (s *Server) registerClient(conn net.Conn) *Client {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := fmt.Sprintf("client-%d", len(s.clients)+1)
	client := &Client{
		id:       id,
		conn:     conn,
		outbound: make(chan string, 10),
		server:   s,
		mode:     ModeLobby,
	}
	s.clients[id] = client

	return client
}

func (s *Server) removeClient(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if client.paired != nil {
		client.paired.paired = nil
		client.paired.mode = ModeLobby
		client.paired.outbound <- "\n[System] Your chat partner has disconnected. You're back in the lobby."
	}
	if client.group != nil {
		client.group.mutex.Lock()
		delete(client.group.members, client.id)
		client.group.mutex.Unlock()
		msg := Message{
			from:    "System",
			content: fmt.Sprintf("%s has left the group", client.name),
			time:    time.Now(),
		}
		client.group.broadcast(msg, client)
	}
	delete(s.clients, client.id)
	delete(s.requests, client.id)
}

func (c *Client) showPrompt() {
	switch c.mode {
	case ModeLobby:
		c.outbound <- "\n=== LOBBY MENU ===\n" +
			"1. /list - Show all clients\n" +
			"2. /groups - List all groups\n" +
			"3. /connect <client-id> - Request private chat\n" +
			"4. /join <group-name> - Join/Create group chat\n" +
			"5. /name <new-name> - Set your display name\n" +
			"6. /quit - Exit application\n" +
			"Current status: In Lobby as " + c.name + " (" + c.id + ")\n" +
			"================\n"
	case ModePrivateChat:
		c.outbound <- "\n=== PRIVATE CHAT ===\n" +
			"1. /disconnect - Return to lobby\n" +
			"2. /quit - Exit application\n" +
			"Currently chatting with: " + c.paired.name + "\n" +
			"================\n"
	case ModeGroupChat:
		c.outbound <- "\n=== GROUP CHAT ===\n" +
			"1. /members - List group members\n" +
			"2. /history - Show recent messages\n" +
			"3. /leave - Return to lobby\n" +
			"4. /quit - Exit application\n" +
			"Current group: " + c.group.name + "\n" +
			"================\n"
	}
}

func (c *Client) handleConnection() {
	defer func() {
		c.server.removeClient(c)
		c.conn.Close()
	}()

	c.name = c.id // Default name is the ID
	c.showPrompt()

	go c.handleOutbound()

	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		msg := scanner.Text()

		if msg == "/quit" {
			return
		}

		if c.handleCommand(msg) {
			c.showPrompt()
			continue
		}

		switch c.mode {
		case ModePrivateChat:
			if c.paired != nil {
				timestamp := time.Now().Format("15:04:05")
				c.paired.outbound <- fmt.Sprintf("[%s] %s: %s", timestamp, c.name, msg)
			}
		case ModeGroupChat:
			if c.group != nil {
				message := Message{
					from:    c.name,
					content: msg,
					time:    time.Now(),
				}
				c.group.broadcast(message, c)
			}
		case ModeLobby:
			c.outbound <- "You're in the lobby. Use /help to see available commands."
		}
	}
}

func (c *Client) handleOutbound() {
	for msg := range c.outbound {
		fmt.Fprintf(c.conn, "%s\n", msg)
	}
}

func (c *Client) handleCommand(msg string) bool {
	parts := strings.Fields(msg)
	if len(parts) == 0 {
		return false
	}

	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "/help":
		c.showPrompt()
		return true

	case "/name":
		if len(args) > 0 {
			c.name = strings.Join(args, " ")
			c.outbound <- fmt.Sprintf("Name changed to: %s", c.name)
			return true
		}

	case "/list":
		if c.mode == ModeLobby {
			var clients []string
			c.server.mutex.RLock()
			for _, client := range c.server.clients {
				status := "🟢 Available"
				if client.mode == ModePrivateChat {
					status = "🔴 In private chat"
				} else if client.mode == ModeGroupChat {
					status = "👥 In group: " + client.group.name
				}
				clients = append(clients, fmt.Sprintf("%s (%s) - %s", client.name, client.id, status))
			}
			c.server.mutex.RUnlock()
			c.outbound <- fmt.Sprintf("Connected clients:\n%s", strings.Join(clients, "\n"))
			return true
		}

	case "/groups":
		if c.mode == ModeLobby {
			var groups []string
			c.server.mutex.RLock()
			for name, group := range c.server.groups {
				groups = append(groups, fmt.Sprintf("%s (%d members)", name, len(group.members)))
			}
			c.server.mutex.RUnlock()
			if len(groups) == 0 {
				c.outbound <- "No active groups"
			} else {
				c.outbound <- fmt.Sprintf("Active groups:\n%s", strings.Join(groups, "\n"))
			}
			return true
		}

	case "/connect":
		if c.mode == ModeLobby && len(args) > 0 {
			c.requestConnection(args[0])
			return true
		}

	case "/accept":
		if c.mode == ModeLobby {
			c.acceptConnection()
			return true
		}

	case "/disconnect":
		if c.mode == ModePrivateChat {
			c.paired.outbound <- fmt.Sprintf("[System] %s has disconnected. You're back in the lobby.", c.name)
			c.paired.mode = ModeLobby
			c.paired.paired = nil
			c.paired = nil
			c.mode = ModeLobby
			c.outbound <- "Disconnected from chat. You're back in the lobby."
			return true
		}

	case "/join":
		if c.mode == ModeLobby && len(args) > 0 {
			groupName := strings.Join(args, " ")
			c.server.mutex.Lock()
			group, exists := c.server.groups[groupName]
			if !exists {
				group = c.server.createGroup(groupName)
				c.outbound <- fmt.Sprintf("Created new group: %s", groupName)
			}
			c.server.mutex.Unlock()

			group.mutex.Lock()
			group.members[c.id] = c
			group.mutex.Unlock()

			c.group = group
			c.mode = ModeGroupChat
			msg := Message{
				from:    "System",
				content: fmt.Sprintf("%s has joined the group", c.name),
				time:    time.Now(),
			}
			group.broadcast(msg, c)
			c.outbound <- fmt.Sprintf("Joined group: %s", groupName)
			return true
		}

	case "/leave":
		if c.mode == ModeGroupChat {
			msg := Message{
				from:    "System",
				content: fmt.Sprintf("%s has left the group", c.name),
				time:    time.Now(),
			}
			c.group.broadcast(msg, c)
			c.group.mutex.Lock()
			delete(c.group.members, c.id)
			c.group.mutex.Unlock()
			c.group = nil
			c.mode = ModeLobby
			c.outbound <- "Left the group. You're back in the lobby."
			return true
		}

	case "/members":
		if c.mode == ModeGroupChat {
			var members []string
			c.group.mutex.RLock()
			for _, member := range c.group.members {
				members = append(members, member.name)
			}
			c.group.mutex.RUnlock()
			c.outbound <- fmt.Sprintf("Group members:\n%s", strings.Join(members, "\n"))
			return true
		}

	case "/history":
		if c.mode == ModeGroupChat {
			c.group.mutex.RLock()
			if len(c.group.messages) == 0 {
				c.outbound <- "No message history"
			} else {
				var history []string
				for _, msg := range c.group.messages {
					history = append(history, fmt.Sprintf("[%s] %s: %s",
						msg.time.Format("15:04:05"),
						msg.from,
						msg.content))
				}
				c.outbound <- fmt.Sprintf("Recent messages:\n%s", strings.Join(history, "\n"))
			}
			c.group.mutex.RUnlock()
			return true
		}
	}

	return false
}

func main() {
	server := NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Enhanced Chat server running on :8080")
	fmt.Println("Features:")
	fmt.Println("- Private chat with /connect and /disconnect")
	fmt.Println("- Group chat with /join and /leave")
	fmt.Println("- Custom display names with /name")
	fmt.Println("- Message history in groups")
	fmt.Println("- Status indicators for users")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		client := server.registerClient(conn)
		go client.handleConnection()
	}
}
