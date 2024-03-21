import { Injectable } from '@angular/core';
import {Sensor} from "../interfaces/sensor";
import {catchError, Observable, of, retry, throwError} from "rxjs";
import {HttpClient, HttpErrorResponse} from "@angular/common/http";
import {BackendURL} from "./app.config";
import { interval as rxjsInterval} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SensorService {

  private mockSensors: Sensor[] = [];

  private knownSensors: Sensor[] = []

  constructor(private http: HttpClient) {
    //this.generateMockSensors(15);
    this.refreshKnownSensors();

    rxjsInterval(30000).subscribe( x => {
      this.refreshKnownSensors();
    })
  }

  private refreshKnownSensors(){
    this.getSensorsAsync().subscribe( sensors => {
      console.log("Refreshed KnownSensors!")
      this.knownSensors = sensors;
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

  public getSensorsNow(): Sensor[]{
    return this.knownSensors;
  }

  public getSensorNameNow(sensorID: string): string | undefined {
    return this.knownSensors.find(sensor => sensor.id === sensorID)?.name
  }



  private generateMockSensors(numberOfSensors: number): void{
    for (let i = 0; i < numberOfSensors; i++) {
      this.mockSensors.push({id: crypto.randomUUID(), name: "name"});
    }
  }
}
