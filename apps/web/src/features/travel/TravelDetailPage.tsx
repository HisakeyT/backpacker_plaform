import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Link, useParams } from "react-router-dom";
import { Box, Button, Chip, Typography } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth";
import { TravelPlanList } from "../travelPlan/TravelPlanList";
import { DeleteTravelDialog } from "./DeleteTravelDialog";
import { getTravel, deleteTravel } from "./repository";
import type { Travel } from "./types";

export const TravelDetailPage = () => {
  const { token } = useAuth();
  const { travelId } = useParams();
  const [travel, setTravel] = useState<Travel | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const navigate = useNavigate();
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [deleteError, setDeleteError] = useState<string | null>(null);

  useEffect(() => {
    if (!token || !travelId) return;

    const fetchTravel = async () => {
      try {
        const data = await getTravel(token, Number(travelId));
        setTravel(data);
      } catch {
        setError("旅行が見つかりませんでした");
      } finally {
        setIsLoading(false);
      }
    };

    fetchTravel();
  }, [token, travelId]);

  const handleOpenDialog = () => {
    setDeleteError(null);
    setIsDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!token || !travel) return;
    setIsDeleting(true);
    setDeleteError(null);
    try {
      await deleteTravel(token, travel.id);
      navigate("/travels");
    } catch {
      setDeleteError("旅行の削除に失敗しました");
      setIsDeleting(false);
    }
  };

  if (isLoading) return <Typography>Loading...</Typography>;
  if (error || !travel) {
    return (
      <Typography color="error">
        {error ?? "旅行が見つかりませんでした"}
      </Typography>
    );
  }

  return (
    <Box className="TravelDetailPage">
      <Button component={Link} to="/travels" sx={{ mb: 2 }}>
        ← 旅行一覧へ
      </Button>

      <Box
        className="TravelDetailPage__header"
        sx={{
          display: "flex",
          alignItems: "flex-start",
          justifyContent: "space-between",
          flexWrap: "wrap",
          gap: 2,
          mb: 1,
        }}
      >
        <Typography
          variant="h4"
          component="h1"
          sx={{
            fontSize: { xs: "1.5rem", sm: "2.125rem" },
            wordBreak: "break-word",
          }}
        >
          {travel.title}
        </Typography>

        <Box sx={{ display: "flex", gap: 1 }}>
          <Button
            component={Link}
            to={`/travels/${travel.id}/edit`}
            variant="contained"
            color="primary"
          >
            編集する
          </Button>
          <Button variant="outlined" color="inherit" onClick={handleOpenDialog}>
            削除する
          </Button>
        </Box>
      </Box>

      <Typography color="text.secondary" gutterBottom>
        {travel.startDate} 〜 {travel.endDate}
      </Typography>
      <Chip
        size="small"
        label={travel.isPublic ? "公開" : "非公開"}
        color={travel.isPublic ? "primary" : "default"}
      />

      <Typography variant="h5" component="h2" sx={{ mt: 4, mb: 2 }}>
        旅のプラン
      </Typography>
      <TravelPlanList travelId={travel.id} />

      <DeleteTravelDialog
        open={isDialogOpen}
        title={travel.title}
        error={deleteError}
        isDeleting={isDeleting}
        onClose={() => setIsDialogOpen(false)}
        onConfirm={handleDelete}
      />
    </Box>
  );
};
