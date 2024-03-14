import { Component } from '@angular/core';
import {SensorService} from "../sensor.service";
import {Sensor} from "../../interfaces/sensor";
import {NgForOf} from "@angular/common";
import {FormsModule} from "@angular/forms";

@Component({
  selector: 'app-sensors',
  standalone: true,
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

    }

    private getSensors(){
      this.sensorService.getSensors().subscribe(sensors => {
        this.sensors = sensors;
        console.log(sensors);
      });
    }

}
