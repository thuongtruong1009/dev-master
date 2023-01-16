import amqplib from 'amqplib';

// const amqpUrl = process.env.AMQP_URL || 'amqp://localhost:5672';
const amqpUrl = 'amqp://localhost:5672';

async function processMessage(msg: any) {
    console.log('Processing message', msg.content.toString());
    // call API here
    await new Promise((resolve) => setTimeout(resolve, 1000));
    console.log('Processed message');
}

(async () => {
    const connection = await amqplib.connect(amqpUrl, 'heartbeat=60');
    const channel = await connection.createChannel();
    channel.prefetch(10);
    const queue = 'user.sign_up_email';
    process.once('SIGINT', async () => {
        console.log('Got signint, closing connection');
        await channel.close();
        await connection.close();
        process.exit(0);
    });

    await channel.assertQueue(queue, { durable: true });
    await channel.consume(
        queue,
        async (msg: any) => {
            console.log('Got message');
            await processMessage(msg);
            await channel.ack(msg);
        },
        {
            noAck: false,
            exclusive: false,
            consumerTag: 'email_consumer',
        },
    );
    console.log('Waiting for messages');
})();
