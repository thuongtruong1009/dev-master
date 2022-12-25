const add = (a: number, b: number) => {
    return a + b;
};

const forcast = {
    date: new Date(),
    weather: 'sunny',
};

const logWeather = ({
    date,
    weather,
}: {
    date: Date;
    weather: string;
}): void => {
    console.log(date);
    console.log(weather);
};

console.log(logWeather(forcast));
