const cron = require('node-cron');

// * * * * *
// -> minutes, hours, day of month, month, and day of the week

cron.schedule('* * * * *', () => {
	console.log("Running a task every minute")
});

cron.schedule('30 9 * * *', ()=> {
	console.log("Running a task every day at 9:03 AM");
});

cron.schedule('0 17 * * 1-5', ()=> {
	console.log("Running a task every weekday(Monday-Friday) at 5 PM");
});

cron.schedule('0 */2 * * *', ()=> {
	console.log("Running a task every 2 hours");
});

cron.schedule('0 0 * * 1', ()=> {
	console.log("Running a task every Monday at midnight");
});