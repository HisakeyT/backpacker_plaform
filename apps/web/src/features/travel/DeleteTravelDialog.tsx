import {
  Alert,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@mui/material";

type DeleteTravelDialogProps = {
  open: boolean;
  title: string;
  error: string | null;
  isDeleting: boolean;
  onClose: () => void;
  onConfirm: () => Promise<void>;
};

export const DeleteTravelDialog = ({
  open,
  title,
  error,
  isDeleting,
  onClose,
  onConfirm,
}: DeleteTravelDialogProps) => {
  return (
    <Dialog
      className="DeleteTravelDialog"
      open={open}
      onClose={isDeleting ? undefined : onClose}
      fullWidth
      maxWidth="xs"
    >
      <DialogTitle>旅行を削除しますか？</DialogTitle>
      <DialogContent>
        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}
        <DialogContentText sx={{ wordBreak: "break-word" }}>
          「{title}」と、その旅のプランがすべて削除されます。この操作は取り消せません。
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose} disabled={isDeleting}>
          キャンセル
        </Button>
        <Button
          onClick={onConfirm}
          color="error"
          variant="contained"
          disabled={isDeleting}
        >
          {isDeleting ? "削除中..." : "削除する"}
        </Button>
      </DialogActions>
    </Dialog>
  );
};
