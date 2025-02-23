import streamDeck, {
  action,
  DialDownEvent,
  DialRotateEvent,
  SingletonAction,
  TouchTapEvent,
  WillAppearEvent,
  WillDisappearEvent
} from "@elgato/streamdeck";
import {generateRGBPNG} from "../../util/generateRGBPNG";
import {globalRGBState} from "../state/global-rgb-state";

@action({ UUID: "com.jacob-powers.nanoleaf-controller.true-color-dial" })
export class TrueColorDial extends SingletonAction<DialSettings> {
  private intervalId: NodeJS.Timeout | null = null;
  private lastRGB: RGBSettings = { r: 0, g: 0, b: 0 }; // Store last known RGB values

  override async onWillAppear(ev: WillAppearEvent<DialSettings>): Promise<void> {
    globalRGBState.subscribe((rgb) => this.updateDial(rgb, ev));

    this.intervalId = setInterval(async () => {
      const rgb = globalRGBState.getColor();

      if (rgb.r !== this.lastRGB.r || rgb.g !== this.lastRGB.g || rgb.b !== this.lastRGB.b) {
        this.lastRGB = { ...rgb };
        await this.updateDial(rgb, ev);
      }
    }, 100);

    if (ev.action.isDial()) {
      ev.action.setFeedbackLayout("$A1");
      return await ev.action.setFeedback({
        "icon": generateRGBPNG(globalRGBState.getColor().r, globalRGBState.getColor().g, globalRGBState.getColor().b)
      });
    }
  }

  override async onWillDisappear(ev: WillDisappearEvent<DialSettings>): Promise<void> {
    if (this.intervalId) {
      clearInterval(this.intervalId);
      this.intervalId = null;
    }
  }

  override async onDialRotate(ev: DialRotateEvent<DialSettings>): Promise<void> {}

  override async onDialDown(ev: DialDownEvent<DialSettings>): Promise<void> {
    const rgb = globalRGBState.getColor();
    await this.updateColor(rgb);
  }

  override async onTouchTap(ev: TouchTapEvent<DialSettings>): Promise<void> {}

  private async updateDial(rgb: RGBSettings, ev: WillAppearEvent<DialSettings>) {
    if (ev.action.isDial()) {
      await ev.action.setFeedback({
        "title": "Press to set",  // No effect if set on WillAppear
        "value": "😉",            // No effect if set on WillAppear
        "icon": generateRGBPNG(globalRGBState.getColor().r, globalRGBState.getColor().g, globalRGBState.getColor().b)
      });
    }
  }

  private async updateColor(rgb: RGBSettings) {
    try {
      const response = await fetch("http://localhost:8080/lights/color", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(rgb)
      });
      if (!response.ok) {
        throw new Error("Failed to send color update");
      }
    } catch (error) {
      streamDeck.logger.error("Error updating color:", error);
    }
  }

}

type DialSettings = {
  value?: number;
};
0