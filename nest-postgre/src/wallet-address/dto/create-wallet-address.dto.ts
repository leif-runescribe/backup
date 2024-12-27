import { IsNotEmpty } from 'class-validator';

export class CreateWalletAddressDto {
  @IsNotEmpty()
  address: string;

  @IsNotEmpty()
  userId: number;
}
