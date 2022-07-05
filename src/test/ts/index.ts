import Car from "./class.js";
const cars: Car[] = [];

for (let i = 0; i < 10; i++) {
  cars.push(new Car("Ford", "Mustang"));
}

console.log(cars);
