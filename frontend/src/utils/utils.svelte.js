export const normalizeDate = (date) => {
  return date ? date.split("T")[0] : null
}
