import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { WalletAddress } from './entities/wallet-address.entity';
import { CreateWalletAddressDto } from './dto/create-wallet-address.dto';
import { User } from 'src/users/entities/user.entity';
@Injectable()
export class WalletAddressService {
  constructor(
    @InjectRepository(WalletAddress)
    private walletAddressRepository: Repository<WalletAddress>,
    @InjectRepository(User)
    private usersRepository: Repository<User>,
  ) {}

  findAll(): Promise<WalletAddress[]> {
    return this.walletAddressRepository.find({ relations: ['user'] });
  }

  async findOne(id: number): Promise<WalletAddress> {
    const walletAddress = await this.walletAddressRepository.findOne({
      where: { id },
      relations: ['user'],
    });
    if (!walletAddress) {
      throw new NotFoundException(`WalletAddress with ID ${id} not found`);
    }
    return walletAddress;
  }

  async create(createWalletAddressDto: CreateWalletAddressDto): Promise<WalletAddress> {
    const user = await this.usersRepository.findOne({
      where: { id: createWalletAddressDto.userId },
    });
    if (!user) {
      throw new NotFoundException(`User with ID ${createWalletAddressDto.userId} not found`);
    }
    const walletAddress = this.walletAddressRepository.create({
      address: createWalletAddressDto.address,
      user: user,
    });
    return this.walletAddressRepository.save(walletAddress);
  }

  async update(id: number, updateWalletAddressDto: CreateWalletAddressDto): Promise<WalletAddress> {
    const user = await this.usersRepository.findOne({
      where: { id: updateWalletAddressDto.userId },
    });
    if (!user) {
      throw new NotFoundException(`User with ID ${updateWalletAddressDto.userId} not found`);
    }
    await this.walletAddressRepository.update(id, {
      address: updateWalletAddressDto.address,
      user: user,
    });
    return this.findOne(id);
  }

  async remove(id: number): Promise<void> {
    const result = await this.walletAddressRepository.delete(id);
    if (result.affected === 0) {
      throw new NotFoundException(`WalletAddress with ID ${id} not found`);
    }
  }
}
