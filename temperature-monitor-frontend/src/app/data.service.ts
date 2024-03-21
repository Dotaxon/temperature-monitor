import { Injectable } from '@angular/core';
import {DataPoint, DataCollection} from "../interfaces/dataInterfaces";
import {DateTime, Duration, DurationLike} from "luxon";

@Injectable({
  providedIn: 'root'
})
export class DataService {

  public static minuteMillis: number = 1000*60;
  public static halfhourMillis : number = DataService.minuteMillis*30;
  public static hourMillis: number = DataService.minuteMillis*60;

  halfhour: DurationLike = { minutes: 30 };

  mockDataCollection : DataCollection[] = [];

  constructor() {
    this.generateMockHourCollections(2);
    console.log(this.mockDataCollection);
  }


























  generateMockHourCollections(numberOfHours: number) : void {
    let date = new Date(2023, 1, 1, 0, 30);
    for (let i = 0; i < numberOfHours; i++) {
      this.mockDataCollection.push(this.generateHour(DateTime.fromJSDate(date).plus(0 * i * DataService.hourMillis)));
    }
  }

  generateHour(start: DateTime): DataCollection{

    let data: DataPoint[] = [];
    for (let i = 0; i < 50; i++) {
      data.push({temp : Math.random() * 30, time: start.plus(i * DataService.minuteMillis).toJSDate() })
    }

    let avg : number = data.map(x => x.temp)
      .reduce((previousValue, currentValue) => currentValue+previousValue, 0) / data.length;

    return {sensorID: "abc", data: data};
  }


}
