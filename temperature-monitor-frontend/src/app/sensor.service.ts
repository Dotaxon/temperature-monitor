import { Injectable } from '@angular/core';
import {Sensor} from "../interfaces/sensor";
import {of} from "rxjs";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class SensorService {

  private mockSensors: Sensor[] = [];

  constructor(private http: HttpClient) {
    this.generateMockSensors(15);
    console.log("constructor sensor service")
  }

  private generateMockSensors(numberOfSensors: number): void{
    for (let i = 0; i < numberOfSensors; i++) {
        this.mockSensors.push({id: crypto.randomUUID(), name: "name"});
    }
  }

  public saveSensors(sensors: Sensor[]){
    this.mockSensors = sensors;
  }

  public getSensors(){
    return this.http.get<Sensor[]>("http://localhost:3000/sensors")
    // return of(this.mockSensors);
  }

  public getSensor(id : string){
    return of(this.mockSensors.find((sensor, index, obj) => sensor.id === id));
  }
}
