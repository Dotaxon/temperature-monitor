import {ApplicationConfig, importProvidersFrom} from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import {HttpClientModule} from "@angular/common/http";

export const appConfig: ApplicationConfig = {
  providers: [provideRouter(routes), importProvidersFrom(HttpClientModule)]
};

//export const BackendURL: string = "https://localhost:3000"
export const BackendURL: string = "https://RPI-Heizung.olbring.org:3000"
