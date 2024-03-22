export interface Sensor{
  id : string,
  name : string
}

export interface SensorWithTemp{
  sensor: Sensor,
  temp: number
}

export interface SelectedSensors{
  name: string
  selected: boolean
}
