import { useState } from "react";
import { Dialog, DialogContent, DialogTitle } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth";
import { createTravelPlan } from "./repository";
import { TravelPlanForm, type TravelPlanFormValue } from "./TravelPlanForm";

type TravelPlanDialogProps = {
  open: boolean;
  travelId: number;
  startDate: string;
  endDate: string;
  // nextSortOrder: (date: string) => number;
  onClose: () => void;
  onSaved: () => void;
};

export const TravelPlanDialog = ({
  open,
  travelId,
  startDate,
  endDate,
  // nextSortOrder,
  onClose,
  onSaved,
}: TravelPlanDialogProps) => {
  const { token } = useAuth();
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

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
      await createTravelPlan(token, travelId, {
        ...value,
        sortOrder: 0,
        // sortOrder: nextSortOrder(value.date),
      });
      onSaved();
      onClose();
    } catch {
      setError("プランの追加に失敗しました");
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
      <DialogTitle>プランを追加</DialogTitle>
      <DialogContent sx={{ pt: 1 }}>
        <TravelPlanForm
          defaultDate={startDate}
          minDate={startDate}
          maxDate={endDate}
          submitLabel="追加する"
          error={error}
          isSubmitting={isSubmitting}
          onCancel={handleClose}
          onSubmit={handleSubmit}
        />
      </DialogContent>
    </Dialog>
  );
};
