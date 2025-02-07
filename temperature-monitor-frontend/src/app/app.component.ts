import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {SensorsComponent} from "./sensors/sensors.component";
import {SidebarComponent} from "./sidebar/sidebar.component";

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, SidebarComponent],
    templateUrl: './app.component.html',
    styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'temperature-monitor-frontend';
}
