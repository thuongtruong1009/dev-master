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
function processMessage(msg) {
    return __awaiter(this, void 0, void 0, function* () {
        console.log('Processing message', msg.content.toString());
        // call API here
        yield new Promise((resolve) => setTimeout(resolve, 1000));
        console.log('Processed message');
    });
}
(() => __awaiter(void 0, void 0, void 0, function* () {
    const connection = yield amqplib_1.default.connect(amqpUrl, 'heartbeat=60');
    const channel = yield connection.createChannel();
    channel.prefetch(10);
    const queue = 'user.sign_up_email';
    process.once('SIGINT', () => __awaiter(void 0, void 0, void 0, function* () {
        console.log('Got signint, closing connection');
        yield channel.close();
        yield connection.close();
        process.exit(0);
    }));
    yield channel.assertQueue(queue, { durable: true });
    yield channel.consume(queue, (msg) => __awaiter(void 0, void 0, void 0, function* () {
        console.log('Got message');
        yield processMessage(msg);
        yield channel.ack(msg);
    }), {
        noAck: false,
        exclusive: false,
        consumerTag: 'email_consumer',
    });
    console.log('Waiting for messages');
}))();
