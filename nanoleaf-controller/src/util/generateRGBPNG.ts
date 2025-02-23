import {createCanvas} from "canvas";

/**
 * Generates a PNG image filled with the specified RGB color and returns it as a Base64 string.
 * @param r - Red value (0-255)
 * @param g - Green value (0-255)
 * @param b - Blue value (0-255)
 * @returns Base64-encoded PNG image
 */
export function generateRGBPNG(r: number, g: number, b: number): string {
  const size = 144; // Image size
  const canvas = createCanvas(size, size);
  const ctx = canvas.getContext("2d");

  ctx.fillStyle = `rgb(${r}, ${g}, ${b})`;
  ctx.fillRect(0, 0, size, size);

  const pngString =  canvas.toBuffer("image/png").toString("base64");
  return `data:image/png;base64,${pngString}`;
}