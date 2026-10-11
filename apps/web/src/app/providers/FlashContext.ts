import { createContext } from "react";
import type { AlertColor } from "@mui/material";

export type FlashContextValue = {
  showFlash: (message: string, severity?: AlertColor) => void;
};

export const FlashContext = createContext<FlashContextValue | null>(null);
