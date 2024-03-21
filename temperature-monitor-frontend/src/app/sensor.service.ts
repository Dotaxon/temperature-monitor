import { Injectable } from '@angular/core';
import {Sensor} from "../interfaces/sensor";
import {catchError, of, retry, throwError} from "rxjs";
import {HttpClient, HttpErrorResponse} from "@angular/common/http";
import {BackendURL} from "./app.config";

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

  private errorHandler(error : HttpErrorResponse){
    if (error.status === 0){
      //Client side or network error no status code
      console.error('An client side or network error ocuured: ', error.error)
    }
    else {
      console.error(`Backend returned code ${error.status}, body was: `, error.error)
    }

    return throwError( () => new Error('Something bad happened; please try again later.'))
  }

  // public saveSensors(sensors: Sensor[]){
  //   this.mockSensors = sensors;
  // }

  public saveSensor(sensor: Sensor){
    this.http.patch(BackendURL + "/sensor/update", sensor)
      .pipe(
        retry(3),
        catchError(this.errorHandler)
      ).subscribe();
  }

  public getSensors(){
    return this.http.get<Sensor[]>(BackendURL + "/sensors")
      .pipe(
        retry(3),
        catchError(this.errorHandler)
      );
    // return of(this.mockSensors);
  }

  // public getSensor(id : string){
  //   return of(this.mockSensors.find((sensor, index, obj) => sensor.id === id));
  // }
}
