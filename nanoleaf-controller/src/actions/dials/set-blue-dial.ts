import {
  action,
  DialDownEvent,
  DialRotateEvent,
  SingletonAction,
  TouchTapEvent,
  WillAppearEvent
} from "@elgato/streamdeck";
import {generateRGBPNG} from "../../util/generateRGBPNG";
import {globalRGBState} from "../state/global-rgb-state";

@action({ UUID: "com.jacob-powers.nanoleaf-controller.set-blue-dial" })
export class SetBlueDial extends SingletonAction<DialSettings> {

  private tickValues: number[] = [1, 5, 10];
  private tickIndex: number = 0;

  override async onWillAppear(ev: WillAppearEvent<DialSettings>): Promise<void> {
    if (ev.action.isDial()) {
      ev.action.setFeedbackLayout("$A1");
      globalRGBState.setColor("b", ev.payload.settings.value ?? 0);
      return ev.action.setFeedback(
        {
          "title": `Dial rotating ${this.tickValues[this.tickIndex]}`,
          "value": ev.payload.settings.value ?? 0,
          "icon": generateRGBPNG(0, 0, ev.payload.settings.value ?? 0)
        }
      )
    }
  }

  override async onDialRotate(ev: DialRotateEvent<DialSettings>): Promise<void> {
    let value = ev.payload.settings.value ?? 0;
    value = value + (ev.payload.ticks * this.tickValues[this.tickIndex]);

    if (value < 0) value = 0;
    if (value > 255) value = 255;

    await ev.action.setSettings({ value });
    globalRGBState.setColor("b", value);
    await ev.action.setFeedback(
      {
        "title": `Dial rotating ${ev.payload.ticks * this.tickValues[this.tickIndex]}`,
        "value": value,
        "icon": generateRGBPNG(0, 0, value)
      }
    );
  }

  override async onDialDown(ev: DialDownEvent<DialSettings>): Promise<void> {
    this.tickIndex = (this.tickIndex + 1) % this.tickValues.length;
    await ev.action.setFeedback({
      "title": ` Tick set to ${this.tickValues[this.tickIndex]}`
    })
  }

  override async onTouchTap(ev: TouchTapEvent<DialSettings>): Promise<void> {
    await ev.action.setSettings({ value: 255 });
    await ev.action.setFeedback({
      "value": 255,
      "icon": generateRGBPNG(0, 0, 255)
    });
  }
}



type DialSettings = {
  value?: number;
};
