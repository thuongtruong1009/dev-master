"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const amqplib_1 = __importDefault(require("amqplib"));
const amqpUrl = process.env.AMQP_URL || 'amqp://localhost:5672';
(() => __awaiter(void 0, void 0, void 0, function* () {
    const connection = yield amqplib_1.default.connect(amqpUrl, 'heartbeat=60');
    const channel = yield connection.createChannel();
    try {
        console.log('Publishing');
        const exchange = 'user.signed_up';
        const queue = 'user.sign_up_email';
        const routingKey = 'sign_up_email';
        yield channel.assertExchange(exchange, 'direct', { durable: true });
        yield channel.assertQueue(queue, { durable: true });
        yield channel.bindQueue(queue, exchange, routingKey);
        const msg = {
            id: Math.floor(Math.random() * 1000),
            email: 'user@domail.com',
            name: 'firstname lastname',
        };
        yield channel.publish(exchange, routingKey, Buffer.from(JSON.stringify(msg)));
        console.log('Published');
    }
    catch (e) {
        console.error('Error in publishing message', e);
    }
    finally {
        yield channel.close();
        yield connection.close();
        console.info('Channel and connection closed');
    }
    process.exit(0);
}))();
