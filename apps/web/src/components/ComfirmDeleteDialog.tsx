import {
  Alert,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@mui/material";

type ConfirmDeleteDialogProps = {
  open: boolean;
  heading: string;
  description: string;
  error: string | null;
  isDeleting: boolean;
  onClose: () => void;
  onConfirm: () => Promise<void>;
};

export const ConfirmDeleteDialog = ({
  open,
  heading,
  description,
  error,
  isDeleting,
  onClose,
  onConfirm,
}: ConfirmDeleteDialogProps) => {
  return (
    <Dialog
      className="ConfirmDeleteDialog"
      open={open}
      onClose={isDeleting ? undefined : onClose}
      fullWidth
      maxWidth="xs"
    >
      <DialogTitle>{heading}</DialogTitle>
      <DialogContent>
        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}
        <DialogContentText sx={{ wordBreak: "break-word" }}>
          {description}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button color="inherit" onClick={onClose} disabled={isDeleting}>
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
