import {ApplicationConfig, importProvidersFrom} from '@angular/core';
import {provideRouter, withHashLocation} from '@angular/router';

import { routes } from './app.routes';
import {} from "@angular/common/http";

export const appConfig: ApplicationConfig = {
  providers: [provideRouter(routes, withHashLocation()), importProvidersFrom(HttpClientModule)]
};

//export const BackendURL: string = "https://localhost:3000"
export const BackendURL: string = "https://RPI-Heizung.olbring.org:3000"
