import { Routes } from '@angular/router';
import {SensorsComponent} from "./sensors/sensors.component";
import {DataComponent} from "./data/data.component";

export const routes: Routes = [
  { path: '', redirectTo: 'data', pathMatch: "full"},
  { path: 'sensors', component: SensorsComponent },
  { path: 'data', component: DataComponent}
];
