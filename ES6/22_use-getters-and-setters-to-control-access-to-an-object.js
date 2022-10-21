class Thermostat {
    constructor(temperature) {
        this._temperature = (temperature - 32) / 1.8;
    }
    get temperature() {
        return this._temperature;
    }
    set temperature(updateTemperature) {
        this._temperature = updateTemperature;
    }
}

const thermos = new Thermostat(76); // Setting in Fahrenheit scale
let temp = thermos.temperature; // 24.44 in Celsius
thermos.temperature = 26;
temp = thermos.temperature; // 26 in Celsius
