import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { Box, Button, Chip, Typography } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth";
import { TravelPlanList } from "../travelPlan/TravelPlanList";
import { getTravel } from "./repository";
import type { Travel } from "./types";

export const TravelDetailPage = () => {
  const { token } = useAuth();
  const { travelId } = useParams();
  const [travel, setTravel] = useState<Travel | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

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
      <Typography variant="h4" component="h1" gutterBottom sx={{
        fontSize: { xs: "1.5rem", sm: "2.125rem" },
        wordBreak: "break-word",
      }}>
        {travel.title}
      </Typography>
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
    </Box>
  );
};
