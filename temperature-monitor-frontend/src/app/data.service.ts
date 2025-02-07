import { Injectable } from '@angular/core';
import {
  DataPoint,
  DataCollection,
  GetDataEntriesRequestBody,
  GetDataEntriesRequestResponse
} from "../interfaces/dataInterfaces";
import {DateTime, Duration, DurationLike} from "luxon";
import {CollectionIntervalEnum} from "../enums/Interval";
import { HttpClient } from "@angular/common/http";
import {BackendURL} from "./app.config";

@Injectable({
  providedIn: 'root'
})
export class DataService {

  public static minuteMillis: number = 1000*60;
  public static halfhourMillis : number = DataService.minuteMillis*30;
  public static hourMillis: number = DataService.minuteMillis*60;

  halfhour: DurationLike = { minutes: 30 };

  mockDataCollections : DataCollection[] = [];

  constructor(private http: HttpClient) {
    this.generateMockHourCollections(2);
    console.log(this.mockDataCollections);
  }

  public getDataEntries(sensorID: string, startTime: DateTime, endTime: DateTime, interval: CollectionIntervalEnum) {
    let requestBody : GetDataEntriesRequestBody = {
      startTime: startTime.toUTC().toUnixInteger(),
      endTime: endTime.toUTC().toUnixInteger(),
      interval: interval,
      sensorID: sensorID
    }

    return this.http.post<GetDataEntriesRequestResponse>(BackendURL + "/data", requestBody)
  }
























  generateMockHourCollections(numberOfHours: number) : void {
    let date = new Date(2023, 1, 1, 0, 30);
    for (let i = 0; i < numberOfHours; i++) {
      this.mockDataCollections.push(this.generateHour(DateTime.fromJSDate(date).plus(0 * i * DataService.hourMillis)));
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
