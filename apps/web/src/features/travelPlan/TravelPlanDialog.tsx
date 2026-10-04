import { useState } from "react";
import { Dialog, DialogContent, DialogTitle } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth";
import { createTravelPlan, updateTravelPlan } from "./repository";
import { TravelPlanForm, type TravelPlanFormValue } from "./TravelPlanForm";
import type { TravelPlan } from "./types";

type TravelPlanDialogProps = {
  open: boolean;
  travelId: number;
  startDate: string;
  endDate: string;
  plan?: TravelPlan | null;
  onClose: () => void;
  onSaved: () => void;
};

export const TravelPlanDialog = ({
  open,
  travelId,
  startDate,
  endDate,
  plan,
  onClose,
  onSaved,
}: TravelPlanDialogProps) => {
  const { token } = useAuth();
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const isEdit = plan != null;

  const handleClose = () => {
    if (isSubmitting) return;
    setError(null);
    onClose();
  };

  const handleSubmit = async (value: TravelPlanFormValue) => {
    if (!token) return;
    setIsSubmitting(true);
    setError(null);
    try {
      if (plan) {
        await updateTravelPlan(token, travelId, plan.id, value);
      } else {
        await createTravelPlan(token, travelId, { ...value, sortOrder: 0 });
      }
      onSaved();
      onClose();
    } catch {
      setError(isEdit ? "プランの更新に失敗しました" : "プランの追加に失敗しました");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <Dialog
      className="TravelPlanDialog"
      open={open}
      onClose={handleClose}
      fullWidth
      maxWidth="sm"
    >
      <DialogTitle>{isEdit ? "プランを編集" : "プランを追加"}</DialogTitle>
      <DialogContent sx={{ pt: 1 }}>
        <TravelPlanForm
          initialValue={
            plan
              ? { date: plan.date, place: plan.place, content: plan.content }
              : undefined
          }
          defaultDate={startDate}
          minDate={startDate}
          maxDate={endDate}
          submitLabel={isEdit ? "保存する" : "追加する"}
          error={error}
          isSubmitting={isSubmitting}
          onCancel={handleClose}
          onSubmit={handleSubmit}
        />
      </DialogContent>
    </Dialog>
  );
};
