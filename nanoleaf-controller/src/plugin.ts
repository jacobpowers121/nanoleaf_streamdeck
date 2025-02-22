import streamDeck, { LogLevel } from "@elgato/streamdeck";

import {DisplayRed} from "./actions/display-red";
import {DisplayColor} from "./actions/display-color";
import {DisplayGreen} from "./actions/display-green";
import {DisplayBlue} from "./actions/display-blue";

streamDeck.logger.setLevel(LogLevel.TRACE);

streamDeck.logger.info("Plugin is starting...");
streamDeck.actions.registerAction(new DisplayRed());
streamDeck.actions.registerAction(new DisplayGreen());
streamDeck.actions.registerAction(new DisplayBlue());
streamDeck.actions.registerAction(new DisplayColor());

streamDeck.connect();