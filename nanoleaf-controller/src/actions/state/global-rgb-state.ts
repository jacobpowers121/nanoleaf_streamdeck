import {EventEmitter} from "events";

class GlobalRGBState extends EventEmitter {
  private rgb: RGBSettings = { r: 0, g: 0, b: 0 };

  setColor(channel: keyof RGBSettings, value: number) {
    this.rgb[channel] = value;
    this.emit("update", this.rgb); // Notify listeners (true color dial)
  }

  getColor(): RGBSettings {
    return this.rgb;
  }

  subscribe(listener: (rgb: RGBSettings) => void) {
    this.on("update", listener);
  }
}

export const globalRGBState = new GlobalRGBState();
