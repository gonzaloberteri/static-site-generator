import Car from "./class";
const cars = [];

for (let i = 0; i < 10; i++) {
  cars.push(new Car("Ford", "Mustang"));
}

console.log(cars);
