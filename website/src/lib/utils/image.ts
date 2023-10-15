export  async function getImgDimensions(file: File) {
  const img: HTMLImageElement = new Image();
  img.src = URL.createObjectURL(file);
  await img.decode();
  return {
    width: img.width,
    height: img.height,
  }
}
