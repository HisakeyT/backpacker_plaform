import { useEffect, useState } from "react";
import { Alert, Snackbar } from "@mui/material";
import { useLocation } from "react-router-dom";
import type { LoginLocationState } from "./types";

export const RegisteredNotice = () => {
  const location = useLocation();
  const state = location.state as LoginLocationState;
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (state?.from === "register") setOpen(true);
  }, [state]);

  return (
    <Snackbar
      open={open}
      autoHideDuration={5000}
      onClose={() => setOpen(false)}
      anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
    >
      <Alert severity="success" variant="filled" onClose={() => setOpen(false)}>
        登録しました。ログインしてください
      </Alert>
    </Snackbar>
  );
};
