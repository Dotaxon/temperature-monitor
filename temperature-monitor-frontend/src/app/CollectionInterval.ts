import {CollectionIntervalEnum} from "../enums/Interval";
import {KeyValue} from "@angular/common";

export class CollectionInterval {

  private currentInterval: CollectionIntervalEnum;

  private static readonly availableValues: KeyValue<string, CollectionIntervalEnum>[] = [
    {key: "Minute", value: CollectionIntervalEnum.Minute},
    {key: "Stunde", value: CollectionIntervalEnum.Hour},
    {key: "Tag", value: CollectionIntervalEnum.Day},
    {key: "Woche", value: CollectionIntervalEnum.Week}
  ]

  constructor(interval: CollectionIntervalEnum) {
    this.currentInterval = interval;
  }

  public convertIntervalToValueFormatString() : string {
//https://canvasjs.com/docs/charts/chart-options/axisx/valueformatstring/
    switch (this.currentInterval) {
      case CollectionIntervalEnum.Minute:
        return "HH:mm";
      case CollectionIntervalEnum.Hour:
        return "HH:mm";
      case CollectionIntervalEnum.Day:
        return "DDD";
      case CollectionIntervalEnum.Week:
        return "D.M"
    }
  }

  public convertIntervalToValueFormatStringToolTip() : string {
//https://canvasjs.com/docs/charts/chart-options/axisx/valueformatstring/
    switch (this.currentInterval) {
      case CollectionIntervalEnum.Minute:
      case CollectionIntervalEnum.Hour:
        return "D.M HH:mm";
      case CollectionIntervalEnum.Day:
        return "D.M.YYYY";
      case CollectionIntervalEnum.Week:
        return "D.M.YYYY"
    }
  }

  public convertIntervalToIntervalTypeString() : string {
//https://canvasjs.com/docs/charts/chart-options/axisx/intervaltype/
    switch (this.currentInterval) {
      case CollectionIntervalEnum.Minute:
        return "minute";
      case CollectionIntervalEnum.Hour:
        return "hour";
      case CollectionIntervalEnum.Day:
        return "day";
      case CollectionIntervalEnum.Week:
        return "day"
    }
  }

  get CurrentInterval(){
    return this.currentInterval
  }
  static get AvailableValues(){
    return CollectionInterval.availableValues;
  }

  set CurrentInterval(i: CollectionIntervalEnum){
    this.currentInterval = i;
  }

}
