import streamDeck, { action, KeyDownEvent, SingletonAction, WillAppearEvent } from "@elgato/streamdeck";
import {createSolidColorDataURL} from "../util/colorImage";

@action({ UUID: "com.jacob-powers.nanoleaf-controller.display-color" })
export class DisplayColor extends SingletonAction<ColorSettings> {
  override async onWillAppear(ev: WillAppearEvent<ColorSettings>): Promise<void> {
    return await this.fetchAndUpdate(ev);
  }

  override async onKeyDown(ev: KeyDownEvent<ColorSettings>): Promise<void> {
    return await this.fetchAndUpdate(ev);
  }

  private async fetchAndUpdate(
    ev: WillAppearEvent<ColorSettings> | KeyDownEvent<ColorSettings>
  ): Promise<void> {
    try {
      const response = await fetch("http://localhost:8080/lights/color");
      const data = (await response.json()) as { r: number; g: number; b: number };
      const imageDataURL = createSolidColorDataURL(data.r, data.g, data.b);

      await ev.action.setImage(imageDataURL);
    } catch (error) {
      streamDeck.logger.error("Error fetching green value:", error);
    }
  }
}

type ColorSettings = {};
