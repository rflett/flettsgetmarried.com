export function scrollToID(elementId: string, offset?: number) {
  const element = document.getElementById(elementId);

  if (offset === void 0) {
    offset = 0;
  }

  if (element != null) {
    const pixelsFromTop = element.getBoundingClientRect().top;

    const scrollToPosition = window.scrollY + pixelsFromTop - 50 - offset;
    window.scrollTo({ behavior: "smooth", top: scrollToPosition });
  }
};
