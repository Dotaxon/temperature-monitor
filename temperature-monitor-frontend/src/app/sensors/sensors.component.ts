import { Component } from '@angular/core';
import {SensorService} from "../sensor.service";
import {Sensor} from "../../interfaces/sensor";
import {NgForOf} from "@angular/common";
import {FormsModule} from "@angular/forms";

@Component({
    selector: 'app-sensors',
    imports: [
        NgForOf,
        FormsModule
    ],
    templateUrl: './sensors.component.html',
    styleUrl: './sensors.component.css'
})
export class SensorsComponent {

    sensors: Sensor[] = [];

    constructor(private sensorService: SensorService) {
    }

    ngOnInit(){
      this.getSensors();
    }

    onSave(){
      this.sensors.forEach((sensor) => {
        this.sensorService.saveSensor(sensor);
      });
    }

    private getSensors(){
      this.sensorService.getSensorsAsync().subscribe(sensors => {
        this.sensors = sensors;
        console.log(sensors);
      });
    }

    protected getTemp(sensorID: string){
      return this.sensorService.getSensorTempNow(sensorID);
    }

}
