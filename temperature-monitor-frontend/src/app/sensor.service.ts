import { Injectable } from '@angular/core';
import {Sensor, SensorWithTemp} from "../interfaces/sensor";
import {catchError, Observable, of, retry, throwError} from "rxjs";
import {HttpClient, HttpErrorResponse} from "@angular/common/http";
import {BackendURL} from "./app.config";
import { interval as rxjsInterval} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SensorService {

  private mockSensors: Sensor[] = [];

  private knownSensors: SensorWithTemp[] = []

  constructor(private http: HttpClient) {
    //this.generateMockSensors(15);
    this.refreshKnownSensors();

    rxjsInterval(30000).subscribe( x => {
      this.refreshKnownSensors();
    })
  }

  private refreshKnownSensors(){
    this.getSensorsAsync().subscribe( sensors => {
      console.log("Received Sensors!");

      this.getSensorsWithTemps(sensors).subscribe(sensorsWithTemps => {
        console.log("Received Sensors with Temp!")
        this.knownSensors = sensorsWithTemps;
      })

    })
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

  public saveSensor(sensor: Sensor){
    this.http.patch(BackendURL + "/sensor/update", sensor)
      .pipe(
        retry(3),
        catchError(this.errorHandler)
      ).subscribe();
  }

  public getSensorsAsync(){
    return this.http.get<Sensor[]>(BackendURL + "/sensors")
      .pipe(
        retry(3),
        catchError(this.errorHandler)
      );
  }

  public getSensorsWithTemps(sensors : Sensor[]) {
    return this.http.post<SensorWithTemp[]>(BackendURL + "/sensors/temps", sensors);
  }

  public getSensorNameNow(sensorID: string): string | undefined {
    return this.knownSensors.find(sensor => sensor.sensor.id === sensorID)?.sensor?.name;
  }

  public getSensorIDNow(sensorName: string): string | undefined {
    return this.knownSensors.find(sensor => sensor.sensor.name === sensorName)?.sensor?.id;
  }

  public getSensorTempNow(sensorID: string): number | undefined {
    return this.knownSensors.find(sensor => sensor.sensor.id === sensorID)?.temp;
  }

  private generateMockSensors(numberOfSensors: number): void{
    for (let i = 0; i < numberOfSensors; i++) {
      this.mockSensors.push({id: crypto.randomUUID(), name: "name"});
    }
  }
}
