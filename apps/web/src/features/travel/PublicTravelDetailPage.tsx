import { useEffect, useState } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import { Box, Button, Typography } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth.ts";
import { PublicTravelPlanList } from "../travelPlan/PublicTravelPlanList";
import { getPublicTravel, copyTravel } from "./repository";
import type { PublicTravel } from "./types";

export const PublicTravelDetailPage = () => {
  const { travelId } = useParams();
  const navigate = useNavigate();
  const { token } = useAuth();
  const [travel, setTravel] = useState<PublicTravel | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isCopying, setIsCopying] = useState(false);
  const [copyError, setCopyError] = useState<string | null>(null);

  useEffect(() => {
    if (!travelId) return;

    const fetchTravel = async () => {
      try {
        const data = await getPublicTravel(Number(travelId));
        setTravel(data);
      } catch {
        setError("旅行が見つかりませんでした");
      } finally {
        setIsLoading(false);
      }
    };

    fetchTravel();
  }, [travelId]);

  const handleCopy = async () => {
    if (!token) {
      navigate("/login");
      return;
    }
    if (!travel) return;

    setIsCopying(true);
    setCopyError(null);
    try {
      const { id } = await copyTravel(token, travel.id);
      navigate(`/travels/${id}`); // 遷移先は実際のルートに合わせる
    } catch {
      setCopyError("コピーに失敗しました。もう一度お試しください");
      setIsCopying(false);
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
    <Box className="PublicTravelDetailPage" sx={{ p: 2 }}>
      <Button component={Link} to="/public/travels" sx={{ mb: 2 }}>
        ← みんなの旅行記へ
      </Button>

      <Typography
        variant="h4"
        component="h1"
        sx={{
          fontSize: { xs: "1.5rem", sm: "2.125rem" },
          wordBreak: "break-word",
          mb: 1,
        }}
      >
        {travel.title}
      </Typography>

      <Typography color="text.secondary" gutterBottom>
        {travel.startDate} 〜 {travel.endDate}
      </Typography>
      <Typography variant="body2" color="text.secondary">
        by {travel.authorNickname}
      </Typography>

      <Button
        variant="contained"
        onClick={handleCopy}
        disabled={isCopying}
        sx={{ mt: 2 }}
      >
        {isCopying ? "コピー中..." : "この旅を参考にする"}
      </Button>
      {copyError && (
        <Typography color="error" variant="body2" sx={{ mt: 1 }}>
          {copyError}
        </Typography>
      )}

      <Typography variant="h5" component="h2" sx={{ mt: 4, mb: 2 }}>
        旅のプラン
      </Typography>
      <PublicTravelPlanList travelId={travel.id} />
    </Box>
  );
};
