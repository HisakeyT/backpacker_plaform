// src/components/Toast.tsx(名前は好みで。FlashMessage / NoticeSnackbar など)
import { Alert, Snackbar } from "@mui/material";
import type { AlertColor } from "@mui/material";
import type { ReactNode } from "react";

type Props = {
  open: boolean;
  onClose: () => void;
  children: ReactNode;
  severity?: AlertColor;
  autoHideDuration?: number;
};

export const FlashMessage = ({
  open,
  onClose,
  children,
  severity = "success",
  autoHideDuration = 5000,
}: Props) => (
  <Snackbar
    open={open}
    autoHideDuration={autoHideDuration}
    onClose={onClose}
    anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
  >
    <Alert severity={severity} variant="filled" onClose={onClose}>
      {children}
    </Alert>
  </Snackbar>
);
