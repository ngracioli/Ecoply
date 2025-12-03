import { ref, onMounted, onUnmounted } from "vue";

export function useIconSize() {
  const isMobile = ref(false);

  const checkMobile = () => {
    isMobile.value = window.innerWidth < 640;
  };

  onMounted(() => {
    checkMobile();
    window.addEventListener("resize", checkMobile);
  });

  onUnmounted(() => {
    window.removeEventListener("resize", checkMobile);
  });

  const iconSize = (mobileSize: number = 18, desktopSize: number = 20) => {
    return isMobile.value ? mobileSize : desktopSize;
  };

  return {
    isMobile,
    iconSize,
  };
}
