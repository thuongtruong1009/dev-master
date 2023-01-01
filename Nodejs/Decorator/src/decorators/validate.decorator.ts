import 'reflect-metadata';
import { EMetadata } from '../enums/metadata.enum';

export function validateDecorator(...keys: string[]) {
    return function (target: any, key: string, desc: PropertyDescriptor) {
        Reflect.defineMetadata(EMetadata.validator, keys, target, key);
    };
}
