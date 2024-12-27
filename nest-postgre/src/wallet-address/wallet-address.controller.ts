import { Controller, Get, Post, Body, Patch, Param, Delete } from '@nestjs/common';
import { WalletAddressService } from './wallet-address.service';
import { CreateWalletAddressDto } from './dto/create-wallet-address.dto';
import { WalletAddress } from './entities/wallet-address.entity';
@Controller('wallet-address')
export class WalletAddressController {
  constructor(private readonly walletAddressService: WalletAddressService) {}

  @Post()
  create(@Body() createWalletAddressDto: CreateWalletAddressDto): Promise<WalletAddress> {
    return this.walletAddressService.create(createWalletAddressDto);
  }

  @Get()
  findAll(): Promise<WalletAddress[]> {
    return this.walletAddressService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string): Promise<WalletAddress> {
    return this.walletAddressService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateWalletAddressDto: CreateWalletAddressDto): Promise<WalletAddress> {
    return this.walletAddressService.update(+id, updateWalletAddressDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string): Promise<void> {
    return this.walletAddressService.remove(+id);
  }
}
