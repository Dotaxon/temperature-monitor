import { Component } from '@angular/core';
import {NgForOf} from "@angular/common";
import {RouterLink} from "@angular/router";
import {SidebarElement} from "../../interfaces/sidebarElement";

@Component({
    selector: 'app-sidebar',
    imports: [
        NgForOf,
        RouterLink
    ],
    templateUrl: './sidebar.component.html',
    styleUrl: './sidebar.component.css'
})
export class SidebarComponent {

  isOpen: boolean = false;

  sidebarElements : SidebarElement[] = [
    { name: "Sensoren", path: "/sensors" },
    { name: "Daten", path: "/data" }
  ];

  onOpen(): void {
    console.log("onOpen")
    this.isOpen = true;
  }

  onClose(): void {
    this.isOpen = false;
  }

}
