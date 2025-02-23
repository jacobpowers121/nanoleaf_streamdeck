import streamDeck, {LogLevel} from "@elgato/streamdeck";

import {DisplayRed} from "./actions/buttons/display-red";
import {DisplayColor} from "./actions/buttons/display-color";
import {DisplayGreen} from "./actions/buttons/display-green";
import {DisplayBlue} from "./actions/buttons/display-blue";
import {SetRedDial} from "./actions/dials/set-red-dial";
import {SetBlueDial} from "./actions/dials/set-blue-dial";
import {SetGreenDial} from "./actions/dials/set-green-dial";
import {TrueColorDial} from "./actions/dials/true-color-dial";

streamDeck.logger.setLevel(LogLevel.TRACE);

streamDeck.logger.info("Plugin is starting...");
streamDeck.actions.registerAction(new SetRedDial());
streamDeck.actions.registerAction(new SetGreenDial());
streamDeck.actions.registerAction(new SetBlueDial());
streamDeck.actions.registerAction(new TrueColorDial());
streamDeck.actions.registerAction(new DisplayRed());
streamDeck.actions.registerAction(new DisplayGreen());
streamDeck.actions.registerAction(new DisplayBlue());
streamDeck.actions.registerAction(new DisplayColor());

streamDeck.connect();