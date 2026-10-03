import { useEffect, useState } from "react";
import { Box, Card, Typography } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth";
import { getTravelPlans } from "./repository";
import type { TravelPlan } from "./types";
import { TravelPlanCard } from "./TravelPlanCard";

type Props = {
  travelId: number;
  reloadKey?: number;
};

const sortPlans = (plans: TravelPlan[]): TravelPlan[] => {
  return [...plans].sort((a, b) => {
    if (a.date !== b.date) return a.date < b.date ? -1 : 1;
    return a.sortOrder - b.sortOrder;
  });
};

const groupByDate = (plans: TravelPlan[]): [string, TravelPlan[]][] => {
  const map = new Map<string, TravelPlan[]>();
  for (const plan of plans) {
    const list = map.get(plan.date) ?? [];
    list.push(plan);
    map.set(plan.date, list);
  }
  return [...map.entries()];
};

export const TravelPlanList = ({ travelId, reloadKey = 0 }: Props) => {
  const { token } = useAuth();
  const [plans, setPlans] = useState<TravelPlan[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!token) return;

    const fetchPlans = async () => {
      try {
        const data = await getTravelPlans(token, travelId);
        setPlans(data);
      } catch {
        setError("予定の取得に失敗しました");
      } finally {
        setIsLoading(false);
      }
    };

    fetchPlans();
  }, [token, travelId, reloadKey]);

  if (isLoading) return <Typography>Loading...</Typography>;
  if (error) return <Typography color="error">{error}</Typography>;
  if (plans.length === 0) {
    return (
      <Typography color="text.secondary">まだ予定がありません</Typography>
    );
  }

  const grouped = groupByDate(sortPlans(plans));

  return (
    <Box className="TravelPlanList" sx={{ display: "grid", gap: 4 }}>
      {grouped.map(([date, dayPlans]) => (
        <Box key={date}>
          <Typography variant="h6" gutterBottom>
            {date}
          </Typography>
          <Box
            sx={{
              display: "grid",
              gap: 2,
              justifyContent: "center",
              gridTemplateColumns: "repeat(auto-fill, minmax(min(100%, 220px), 1fr))",
            }}
          >
            {dayPlans.map((plan) => (
              <TravelPlanCard key={plan.id} plan={plan} />
            ))}
          </Box>
        </Box>
      ))}
    </Box>
  );
};
