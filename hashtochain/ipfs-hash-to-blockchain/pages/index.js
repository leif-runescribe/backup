import React, { useState, useEffect } from "react";

function App() {
  const [file, setFile] = useState(null);
  const [fileName, setFileName] = useState("");
  const [result, setResult] = useState("");

  useEffect(() => {}, []);

  const handleChange = (e) => {
    if (e.target.name === "filename") {
      setFileName(e.target.value);
    }
    if (e.target.name === "file") {
      setFile(e.target.files[0]);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      var formData = new FormData();
      formData.append("filename", fileName);
      formData.append("file", file);

      const res = await fetch("/api/uploadData", {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        throw new Error("Network response is not ok");
      }
      const data = await res.json();
      setResult(data.message);
    } catch (err) {
      console.error(err);
    }
  };

  return (
      <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            height: "100vh",
            background: "lightblue", // Set the outer background to light blue
          }}
      >
        <img
            src={"/logo.png"}
            alt="Logo"
            style={{
              width: "150px", // Adjust the size of your logo as needed
              marginBottom: "20px",
            }}
        />
        <form
            style={{
              background: "white",
              padding: "20px",
              borderRadius: "5px", // Rounded edges
              boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
            }}
            onSubmit={handleSubmit}
        >
          <h1>Store Hash to the Blockchain</h1>
          <label>Enter Unique Filename:</label>
          <input
              type="text"
              name="filename"
              value={fileName}
              onChange={handleChange}
              style={{
                width: "100%",
                padding: "10px",
                margin: "5px 0",
                borderRadius: "3px",
                border: "1px solid #ccc",
              }}
          />
          <input
              type="file"
              name="file"
              onChange={handleChange}
              style={{
                width: "100%",
                padding: "10px",
                margin: "5px 0",
                borderRadius: "3px",
                border: "1px solid #ccc",
              }}
          />
          <input
              type="Submit"
              style={{
                width: "100%",
                padding: "10px",
                margin: "10px 0",
                background: "#007bff",
                color: "white",
                borderRadius: "3px",
                border: "none",
                cursor: "pointer",
              }}
              value="Submit"
          />
          {result && <p>{result}</p>}
        </form>
      </div>
  );
}

export default App;
