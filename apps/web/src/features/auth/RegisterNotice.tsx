import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import type { LoginLocationState } from "./types";
import { FlashMessage } from "../../components/FlashMessage";

export const RegisteredNotice = () => {
  const location = useLocation();
  const state = location.state as LoginLocationState;
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (state?.from === "register") setOpen(true);
  }, [state]);

  return (
    <FlashMessage open={open} onClose={() => setOpen(false)} severity="success">
      登録が完了しました。ログインしてください。
    </FlashMessage>
  );
};
