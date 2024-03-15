import { Injectable } from '@angular/core';
import {DataPoint, hourCollection} from "../interfaces/dataInterfaces";
import {reduce} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class DataService {

  minute: number = 1000*60;
  halfhour : number = this.minute*30;

  mockHourCollection : hourCollection[] = [];

  constructor() {
    this.generateMockHourCollections(24);
    console.log(this.mockHourCollection);
  }

  generateMockHourCollections(numberOfHours: number) : void{
    for (let i = 0; i < numberOfHours; i++) {
      this.mockHourCollection.push(this.generateHour(new Date(2023, 1, 1, i, 30)));
    }
  }

  generateHour(start: Date): hourCollection{
    let data: DataPoint[] = [];
    for (let i = 0; i < 60; i++) {
      data.push({temp : Math.random(), time: new Date(start.getTime() + i * this.minute)})
    }

    let avg : number = data.map(x => x.temp)
      .reduce((previousValue, currentValue) => currentValue+previousValue, 0) / data.length;

    return {hour: new Date(start.getTime() + this.halfhour), average: avg, data: data};
  }


}
