import { Injectable } from '@angular/core';
import {DataPoint, dataCollection} from "../interfaces/dataInterfaces";
import {DateTime, Duration, DurationLike} from "luxon";

@Injectable({
  providedIn: 'root'
})
export class DataService {

  public static minuteMillis: number = 1000*60;
  public static halfhourMillis : number = DataService.minuteMillis*30;
  public static hourMillis: number = DataService.minuteMillis*60;

  halfhour: DurationLike = { minutes: 30 };

  mockHourCollection : dataCollection[] = [];

  constructor() {
    this.generateMockHourCollections(24);
    console.log(this.mockHourCollection);
  }

  generateMockHourCollections(numberOfHours: number) : void {
    let date = new Date(2023, 1, 1, 0, 30);
    for (let i = 0; i < numberOfHours; i++) {
      this.mockHourCollection.push(this.generateHour(DateTime.fromJSDate(date).plus(i * DataService.hourMillis)));
    }
  }

  generateHour(start: DateTime): dataCollection{

    let data: DataPoint[] = [];
    for (let i = 0; i < 60; i++) {
      data.push({temp : Math.random(), time: start.plus(i * DataService.minuteMillis).toJSDate() })
    }

    let avg : number = data.map(x => x.temp)
      .reduce((previousValue, currentValue) => currentValue+previousValue, 0) / data.length;

    return {time: new Date(start.plus(this.halfhour).toJSDate()), average: avg, data: data};
  }


}
