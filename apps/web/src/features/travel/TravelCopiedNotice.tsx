import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { FlashMessage } from "../../components/FlashMessage";

type CopiedLocationState = { copied?: boolean } | null;

export const TravelCopiedNotice = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const state = location.state as CopiedLocationState;

  const [open, setOpen] = useState(Boolean(state?.copied));

  // リロードで再表示されないよう、state を消す
  useEffect(() => {
    if (state?.copied) {
      navigate(location.pathname, { replace: true, state: null });
    }
  }, [state, location.pathname, navigate]);

  return (
    <FlashMessage open={open} onClose={() => setOpen(false)}>
      旅行をコピーしました。日付を確認して、あなたの予定に合わせて編集しましょう。
    </FlashMessage>
  );
};
