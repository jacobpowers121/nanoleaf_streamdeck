import {createCanvas} from "canvas"; // Use node-canvas

export function createSolidColorDataURL(r: number, g: number, b: number): string {
  const canvas = createCanvas(72, 72);
  const ctx = canvas.getContext("2d");
  if (ctx) {
    ctx.fillStyle = `rgb(${r}, ${g}, ${b})`;
    ctx.fillRect(0, 0, canvas.width, canvas.height);
  }
  return canvas.toDataURL();
}
