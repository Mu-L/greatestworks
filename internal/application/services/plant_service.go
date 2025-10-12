package services

import (
	"context"
	"time"

	"greatestworks/internal/domain/scene/plant"
)

// PlantService 种植应用服务
type PlantService struct {
	farmRepo plant.FarmRepository
	cropRepo plant.CropRepository
	// seedRepo       plant.SeedRepository // TODO: Define SeedRepository
	harvestRepo plant.HarvestRepository
	// statisticsRepo plant.StatisticsRepository // TODO: Define StatisticsRepository
	cacheRepo    plant.PlantCacheRepository
	plantService *plant.PlantService
}

// NewPlantService 创建种植应用服务
func NewPlantService(
	farmRepo plant.FarmRepository,
	cropRepo plant.CropRepository,
	// seedRepo plant.SeedRepository,
	harvestRepo plant.HarvestRepository,
	// statisticsRepo plant.StatisticsRepository,
	cacheRepo plant.PlantCacheRepository,
	plantService *plant.PlantService,
) *PlantService {
	return &PlantService{
		farmRepo: farmRepo,
		cropRepo: cropRepo,
		// seedRepo:       seedRepo,
		harvestRepo: harvestRepo,
		// statisticsRepo: statisticsRepo,
		cacheRepo:    cacheRepo,
		plantService: plantService,
	}
}

// GetFarmInfo 获取农场信息
func (s *PlantService) GetFarmInfo(ctx context.Context, playerID string) (*FarmDTO, error) {
	// 先从缓存获取
	// TODO: 修复GetFarm方法调用
	// cachedFarm, err := s.cacheRepo.GetFarm(playerID)
	// if err == nil && cachedFarm != nil {
	// 	return s.buildFarmDTO(cachedFarm), nil
	// }

	// 从数据库获取
	// TODO: 修复FindByPlayer方法调用
	// farm, err := s.farmRepo.FindByPlayer(playerID)
	farm := &plant.FarmAggregate{}
	// TODO: 修复err变量
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get farm info: %w", err)
	// }

	// 更新缓存
	// TODO: 修复SetFarm方法调用
	// if err := s.cacheRepo.SetFarm(playerID, farm, time.Hour); err != nil {
	// 	// 缓存更新失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return s.buildFarmDTO(farm), nil
}

// PlantSeed 种植种子
func (s *PlantService) PlantSeed(ctx context.Context, playerID string, plotID string, seedID string) error {
	// 获取农场信息
	// TODO: 修复FindByPlayer方法调用
	// farm, err := s.farmRepo.FindByPlayer(playerID)
	// if err != nil {
	// 	return fmt.Errorf("failed to get farm info: %w", err)
	// }
	// farm := &plant.FarmAggregate{}

	// 获取种子信息
	// TODO: 修复seedRepo字段
	// seed, err := s.seedRepo.FindByID(seedID)
	// if err != nil {
	// 	return fmt.Errorf("failed to get seed info: %w", err)
	// }
	// seed := &plant.Seed{}

	// 种植种子
	// TODO: 修复PlantSeed方法调用
	// crop, err := s.plantService.PlantSeed(farm, plotID, seed)
	// crop := &plant.Crop{}
	// TODO: 修复err变量
	// if err != nil {
	// 	return fmt.Errorf("failed to plant seed: %w", err)
	// }

	// 保存作物
	// TODO: 修复Save方法调用
	// if err := s.cropRepo.Save(crop); err != nil {
	// 	return fmt.Errorf("failed to save crop: %w", err)
	// }

	// 更新农场
	// TODO: 修复Update方法调用
	// if err := s.farmRepo.Update(farm); err != nil {
	// 	return fmt.Errorf("failed to update farm: %w", err)
	// }

	// 更新统计数据
	// TODO: 修复updatePlantingStatistics方法调用
	// if err := s.updatePlantingStatistics(ctx, playerID, seedID); err != nil {
	// 	// 统计更新失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	// 清除缓存
	// if err := s.cacheRepo.DeleteFarm(playerID); err != nil {
	// 	// 缓存清除失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return nil
}

// WaterCrop 浇水
func (s *PlantService) WaterCrop(ctx context.Context, playerID string, cropID string) error {
	// 获取作物信息
	// TODO: 修复FindByID方法调用
	// crop, err := s.cropRepo.FindByID(cropID)
	// if err != nil {
	// 	return fmt.Errorf("failed to get crop info: %w", err)
	// }
	// crop := &plant.Crop{}

	// TODO: 修复GetPlayerID方法调用
	// if crop.GetPlayerID() != playerID {
	// 	return plant.ErrUnauthorized
	// }

	// 浇水
	// TODO: 修复Water方法调用
	// if err := crop.Water(); err != nil {
	// 	return fmt.Errorf("failed to water crop: %w", err)
	// }

	// 更新作物
	// TODO: 修复Update方法调用
	// if err := s.cropRepo.Update(crop); err != nil {
	// 	return fmt.Errorf("failed to update crop: %w", err)
	// }

	// 清除相关缓存
	// TODO: 修复DeleteFarm方法调用
	// if err := s.cacheRepo.DeleteFarm(playerID); err != nil {
	// 	// 缓存清除失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return nil
}

// FertilizeCrop 施肥
func (s *PlantService) FertilizeCrop(ctx context.Context, playerID string, cropID string, fertilizerType plant.FertilizerType) error {
	// 获取作物信息
	// TODO: 修复FindByID方法调用
	// crop, err := s.cropRepo.FindByID(cropID)
	// crop := &plant.Crop{}
	// TODO: 修复err变量
	// if err != nil {
	// 	return fmt.Errorf("failed to get crop info: %w", err)
	// }

	// TODO: 修复GetPlayerID方法调用
	// if crop.GetPlayerID() != playerID {
	// 	return plant.ErrUnauthorized
	// }

	// 施肥
	// TODO: 修复Fertilize方法调用
	// if err := crop.Fertilize(fertilizerType); err != nil {
	// 	return fmt.Errorf("failed to fertilize crop: %w", err)
	// }

	// 更新作物
	// TODO: 修复Update方法调用
	// if err := s.cropRepo.Update(crop); err != nil {
	// 	return fmt.Errorf("failed to update crop: %w", err)
	// }

	// 清除相关缓存
	// TODO: 修复DeleteFarm方法调用
	// if err := s.cacheRepo.DeleteFarm(playerID); err != nil {
	// 	// 缓存清除失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return nil
}

// HarvestCrop 收获作物
func (s *PlantService) HarvestCrop(ctx context.Context, playerID string, cropID string) (*HarvestResultDTO, error) {
	// 获取作物信息
	// TODO: 修复FindByID方法调用
	// crop, err := s.cropRepo.FindByID(cropID)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get crop info: %w", err)
	// }
	// crop := &plant.Crop{}

	// TODO: 修复GetPlayerID方法调用
	// if crop.GetPlayerID() != playerID {
	// 	return nil, plant.ErrUnauthorized
	// }

	// 收获作物
	// TODO: 修复HarvestCrop方法调用
	// harvestResult, err := s.plantService.HarvestCrop(crop)
	harvestResult := &plant.HarvestResult{}
	// TODO: 修复err变量
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to harvest crop: %w", err)
	// }

	// 保存收获记录
	// TODO: 修复Save方法调用
	// if err := s.harvestRepo.Save(harvestResult); err != nil {
	// 	return nil, fmt.Errorf("failed to save harvest record: %w", err)
	// }

	// 删除作物（已收获）
	// TODO: 修复Delete方法调用
	// if err := s.cropRepo.Delete(cropID); err != nil {
	// 	return nil, fmt.Errorf("failed to delete harvested crop: %w", err)
	// }

	// 更新统计数据
	// TODO: 修复updateHarvestStatistics方法调用
	// if err := s.updateHarvestStatistics(ctx, playerID, harvestResult); err != nil {
	// 	// 统计更新失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	// 清除相关缓存
	// TODO: 修复DeleteFarm方法调用
	// if err := s.cacheRepo.DeleteFarm(playerID); err != nil {
	// 	// 缓存清除失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return s.buildHarvestResultDTO(harvestResult), nil
}

// GetCropInfo 获取作物信息
func (s *PlantService) GetCropInfo(ctx context.Context, playerID string, cropID string) (*CropDTO, error) {
	// 获取作物信息
	// TODO: 修复FindByID方法调用
	// crop, err := s.cropRepo.FindByID(cropID)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get crop info: %w", err)
	// }
	crop := &plant.Crop{}

	// TODO: 修复GetPlayerID方法调用
	// if crop.GetPlayerID() != playerID {
	// 	return nil, plant.ErrUnauthorized
	// }

	return s.buildCropDTO(crop), nil
}

// GetPlayerCrops 获取玩家所有作物
func (s *PlantService) GetPlayerCrops(ctx context.Context, playerID string) ([]*CropDTO, error) {
	// 获取玩家所有作物
	// TODO: 修复FindByPlayer方法调用
	// crops, err := s.cropRepo.FindByPlayer(playerID)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get player crops: %w", err)
	// }
	crops := []*plant.Crop{}

	return s.buildCropDTOs(crops), nil
}

// GetAvailableSeeds 获取可用种子
func (s *PlantService) GetAvailableSeeds(ctx context.Context, playerID string) ([]*SeedDTO, error) {
	// 先从缓存获取
	// TODO: 修复GetAvailableSeeds方法调用
	// cachedSeeds, err := s.cacheRepo.GetAvailableSeeds(playerID)
	// if err == nil && len(cachedSeeds) > 0 {
	// 	return s.buildSeedDTOs(cachedSeeds), nil
	// }

	// 从数据库获取
	// TODO: 修复FindAvailableForPlayer方法调用
	// seeds, err := s.seedRepo.FindAvailableForPlayer(playerID)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get available seeds: %w", err)
	// }
	// seeds := []interface{}{} // TODO: 修复plant.Seed类型

	// 更新缓存
	// TODO: 修复SetAvailableSeeds方法调用
	// if err := s.cacheRepo.SetAvailableSeeds(playerID, seeds, time.Hour*2); err != nil {
	// 	// 缓存更新失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return s.buildSeedDTOs([]plant.SeedType{}), nil // TODO: 修复seeds类型
}

// GetHarvestHistory 获取收获历史
func (s *PlantService) GetHarvestHistory(ctx context.Context, playerID string, limit int) ([]*HarvestHistoryDTO, error) {
	// TODO: 修复FindByPlayer方法调用
	// history, err := s.harvestRepo.FindByPlayer(playerID, limit)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get harvest history: %w", err)
	// }
	history := []*plant.HarvestResult{}

	return s.buildHarvestHistoryDTOs(history), nil
}

// GetPlantingStatistics 获取种植统计
func (s *PlantService) GetPlantingStatistics(ctx context.Context, playerID string) (*PlantingStatisticsDTO, error) {
	// TODO: 修复statisticsRepo字段
	// stats, err := s.statisticsRepo.FindByPlayer(playerID)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get planting statistics: %w", err)
	// }
	// stats := &struct{}{} // TODO: 修复plant.PlantingStatistics类型

	return s.buildStatisticsDTO(&plant.FarmStatistics{}), nil // TODO: 修复stats类型
}

// UpdateCropGrowth 更新作物成长（系统调用）
func (s *PlantService) UpdateCropGrowth(ctx context.Context) error {
	// 获取所有需要更新的作物
	// TODO: 修复FindGrowingCrops方法调用
	// crops, err := s.cropRepo.FindGrowingCrops()
	// if err != nil {
	// 	return fmt.Errorf("failed to get growing crops: %w", err)
	// }
	crops := []*plant.Crop{}

	for range crops {
		// 更新作物成长
		// TODO: 修复UpdateCropGrowth方法调用
		// if err := s.plantService.UpdateCropGrowth(crop); err != nil {
		// 	// 单个作物更新失败不影响其他作物
		// 	// TODO: 添加日志记录
		// 	continue
		// }

		// 保存更新后的作物
		// TODO: 修复Update方法调用
		// if err := s.cropRepo.Update(crop); err != nil {
		// 	// TODO: 添加日志记录
		// 	continue
		// }
	}

	return nil
}

// UpgradeFarmPlot 升级农场地块
func (s *PlantService) UpgradeFarmPlot(ctx context.Context, playerID string, plotID string) error {
	// 获取农场信息
	// TODO: 修复FindByPlayer方法调用
	// farm, err := s.farmRepo.FindByPlayer(playerID)
	// if err != nil {
	// 	return fmt.Errorf("failed to get farm info: %w", err)
	// }
	// farm := &plant.FarmAggregate{}

	// 升级地块
	// TODO: 修复UpgradePlot方法调用
	// if err := farm.UpgradePlot(plotID); err != nil {
	// 	return fmt.Errorf("failed to upgrade plot: %w", err)
	// }

	// 更新农场
	// TODO: 修复Update方法调用
	// if err := s.farmRepo.Update(farm); err != nil {
	// 	return fmt.Errorf("failed to update farm: %w", err)
	// }

	// 清除缓存
	// TODO: 修复DeleteFarm方法调用
	// if err := s.cacheRepo.DeleteFarm(playerID); err != nil {
	// 	// 缓存清除失败不影响主流程
	// 	// TODO: 添加日志记录
	// }

	return nil
}

// 私有方法

// updatePlantingStatistics 更新种植统计
func (s *PlantService) updatePlantingStatistics(ctx context.Context, playerID string, seedID string) error {
	// TODO: 修复statisticsRepo字段
	// stats, err := s.statisticsRepo.FindByPlayer(playerID)
	// if err != nil && !plant.IsNotFoundError(err) {
	// 	return err
	// }

	// if stats == nil {
	// 	stats = plant.NewPlantingStatistics(playerID)
	// }
	// stats := &plant.PlantingStatistics{}

	// 更新统计数据
	// TODO: 修复AddPlantedSeed方法调用
	// stats.AddPlantedSeed(seedID)
	// stats.UpdateLastPlantTime(time.Now())

	// 保存统计数据
	// TODO: 修复Save方法调用
	// return s.statisticsRepo.Save(stats)
	return nil
}

// updateHarvestStatistics 更新收获统计
func (s *PlantService) updateHarvestStatistics(ctx context.Context, playerID string, harvestResult *plant.HarvestResult) error {
	// TODO: 修复statisticsRepo字段
	// stats, err := s.statisticsRepo.FindByPlayer(playerID)
	// if err != nil && !plant.IsNotFoundError(err) {
	// 	return err
	// }

	// if stats == nil {
	// 	stats = plant.NewPlantingStatistics(playerID)
	// }
	// stats := &plant.PlantingStatistics{}

	// 更新统计数据
	// TODO: 修复AddHarvestResult方法调用
	// stats.AddHarvestResult(harvestResult.GetCropType(), harvestResult.GetQuantity(), harvestResult.GetQuality())
	// stats.UpdateLastHarvestTime(harvestResult.GetHarvestTime())

	// 保存统计数据
	// TODO: 修复Save方法调用
	// return s.statisticsRepo.Save(stats)
	return nil
}

// buildFarmDTO 构建农场DTO
func (s *PlantService) buildFarmDTO(farm *plant.FarmAggregate) *FarmDTO {
	plots := make([]*PlotDTO, 0)
	for range farm.GetPlots() {
		plots = append(plots, &PlotDTO{
			ID:         "",    // TODO: plot.GetID(),
			Level:      0,     // TODO: plot.GetLevel(),
			SoilType:   "",    // TODO: string(plot.GetSoilType()),
			Fertility:  0,     // TODO: plot.GetFertility(),
			Moisture:   0,     // TODO: plot.GetMoisture(),
			IsOccupied: false, // TODO: plot.IsOccupied(),
			CropID:     "",    // TODO: plot.GetCropID(),
		})
	}

	return &FarmDTO{
		PlayerID:   "", // TODO: farm.GetPlayerID(),
		Level:      0,  // TODO: farm.GetLevel(),
		Experience: 0,  // TODO: farm.GetExperience(),
		Plots:      plots,
		TotalPlots: 0,          // TODO: farm.GetTotalPlots(),
		UsedPlots:  0,          // TODO: farm.GetUsedPlots(),
		CreatedAt:  time.Now(), // TODO: farm.GetCreatedAt(),
		UpdatedAt:  time.Now(), // TODO: farm.GetUpdatedAt(),
	}
}

// buildCropDTO 构建作物DTO
func (s *PlantService) buildCropDTO(crop *plant.Crop) *CropDTO {
	return &CropDTO{
		ID:                   crop.GetID(),
		PlayerID:             crop.GetPlayerID(),
		PlotID:               crop.GetPlotID(),
		SeedID:               crop.GetSeedID(),
		CropType:             string(crop.GetCropType()),
		CurrentStage:         string(crop.GetCurrentStage()),
		GrowthProgress:       crop.GetGrowthProgress(),
		Health:               0,          // TODO: crop.GetHealth(),
		Moisture:             0,          // TODO: crop.GetMoisture(),
		Nutrition:            0,          // TODO: crop.GetNutrition(),
		Quality:              0,          // TODO: crop.GetQuality(),
		PlantedAt:            time.Now(), // TODO: crop.GetPlantedAt(),
		LastWatered:          time.Now(), // TODO: crop.GetLastWatered(),
		LastFertilized:       time.Now(), // TODO: crop.GetLastFertilized(),
		EstimatedHarvestTime: time.Now(), // TODO: crop.GetEstimatedHarvestTime(),
		IsReadyToHarvest:     false,      // TODO: crop.IsReadyToHarvest(),
	}
}

// buildCropDTOs 构建作物DTO列表
func (s *PlantService) buildCropDTOs(crops []*plant.Crop) []*CropDTO {
	dtos := make([]*CropDTO, len(crops))
	for i, crop := range crops {
		dtos[i] = s.buildCropDTO(crop)
	}
	return dtos
}

// buildSeedDTOs 构建种子DTO列表
func (s *PlantService) buildSeedDTOs(seeds []plant.SeedType) []*SeedDTO {
	dtos := make([]*SeedDTO, len(seeds))
	for i, _ := range seeds {
		dtos[i] = &SeedDTO{
			ID:            "", // TODO: seed.GetID(),
			Name:          "", // TODO: seed.GetName(),
			CropType:      "", // TODO: string(seed.GetCropType()),
			GrowthTime:    0,  // TODO: seed.GetGrowthTime(),
			RequiredLevel: 0,  // TODO: seed.GetRequiredLevel(),
			Price:         0,  // TODO: seed.GetPrice(),
			Yield:         0,  // TODO: seed.GetYield(),
			Quality:       0,  // TODO: seed.GetQuality(),
			Description:   "", // TODO: seed.GetDescription(),
		}
	}
	return dtos
}

// buildHarvestResultDTO 构建收获结果DTO
func (s *PlantService) buildHarvestResultDTO(harvestResult *plant.HarvestResult) *HarvestResultDTO {
	return &HarvestResultDTO{
		CropID:      "",               // TODO: harvestResult.GetCropID(),
		CropType:    "",               // TODO: string(harvestResult.GetCropType()),
		Quantity:    0,                // TODO: harvestResult.GetQuantity(),
		Quality:     0,                // TODO: harvestResult.GetQuality(),
		Experience:  0,                // TODO: harvestResult.GetExperience(),
		HarvestTime: time.Now(),       // TODO: harvestResult.GetHarvestTime(),
		Items:       map[string]int{}, // TODO: harvestResult.GetItems(),
	}
}

// buildHarvestHistoryDTOs 构建收获历史DTO列表
func (s *PlantService) buildHarvestHistoryDTOs(history []*plant.HarvestResult) []*HarvestHistoryDTO {
	dtos := make([]*HarvestHistoryDTO, len(history))
	for i, _ := range history {
		dtos[i] = &HarvestHistoryDTO{
			ID:          "",         // TODO: record.GetID(),
			CropType:    "",         // TODO: string(record.GetCropType()),
			Quantity:    0,          // TODO: record.GetQuantity(),
			Quality:     0,          // TODO: record.GetQuality(),
			Experience:  0,          // TODO: record.GetExperience(),
			HarvestTime: time.Now(), // TODO: record.GetHarvestTime(),
		}
	}
	return dtos
}

// buildStatisticsDTO 构建统计DTO
func (s *PlantService) buildStatisticsDTO(stats *plant.FarmStatistics) *PlantingStatisticsDTO {
	return &PlantingStatisticsDTO{
		PlayerID:        "",                          // TODO: stats.GetPlayerID(),
		TotalPlanted:    0,                           // TODO: stats.GetTotalPlanted(),
		TotalHarvested:  0,                           // TODO: stats.GetTotalHarvested(),
		TotalExperience: 0,                           // TODO: stats.GetTotalExperience(),
		CropTypeStats:   map[string]*CropTypeStats{}, // TODO: stats.GetCropTypeStats(),
		AverageQuality:  0,                           // TODO: stats.GetAverageQuality(),
		FavoriteCrop:    "",                          // TODO: string(stats.GetFavoriteCrop()),
		LastPlantTime:   time.Now(),                  // TODO: stats.GetLastPlantTime(),
		LastHarvestTime: time.Now(),                  // TODO: stats.GetLastHarvestTime(),
	}
}

// DTO 定义

// FarmDTO 农场DTO
type FarmDTO struct {
	PlayerID   string     `json:"player_id"`
	Level      int        `json:"level"`
	Experience int64      `json:"experience"`
	Plots      []*PlotDTO `json:"plots"`
	TotalPlots int        `json:"total_plots"`
	UsedPlots  int        `json:"used_plots"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// PlotDTO 地块DTO
type PlotDTO struct {
	ID         string  `json:"id"`
	Level      int     `json:"level"`
	SoilType   string  `json:"soil_type"`
	Fertility  float64 `json:"fertility"`
	Moisture   float64 `json:"moisture"`
	IsOccupied bool    `json:"is_occupied"`
	CropID     string  `json:"crop_id,omitempty"`
}

// CropDTO 作物DTO
type CropDTO struct {
	ID                   string    `json:"id"`
	PlayerID             string    `json:"player_id"`
	PlotID               string    `json:"plot_id"`
	SeedID               string    `json:"seed_id"`
	CropType             string    `json:"crop_type"`
	CurrentStage         string    `json:"current_stage"`
	GrowthProgress       float64   `json:"growth_progress"`
	Health               float64   `json:"health"`
	Moisture             float64   `json:"moisture"`
	Nutrition            float64   `json:"nutrition"`
	Quality              float64   `json:"quality"`
	PlantedAt            time.Time `json:"planted_at"`
	LastWatered          time.Time `json:"last_watered"`
	LastFertilized       time.Time `json:"last_fertilized"`
	EstimatedHarvestTime time.Time `json:"estimated_harvest_time"`
	IsReadyToHarvest     bool      `json:"is_ready_to_harvest"`
}

// SeedDTO 种子DTO
type SeedDTO struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	CropType      string        `json:"crop_type"`
	GrowthTime    time.Duration `json:"growth_time"`
	RequiredLevel int           `json:"required_level"`
	Price         int64         `json:"price"`
	Yield         int           `json:"yield"`
	Quality       float64       `json:"quality"`
	Description   string        `json:"description"`
}

// HarvestResultDTO 收获结果DTO
type HarvestResultDTO struct {
	CropID      string         `json:"crop_id"`
	CropType    string         `json:"crop_type"`
	Quantity    int            `json:"quantity"`
	Quality     float64        `json:"quality"`
	Experience  int64          `json:"experience"`
	HarvestTime time.Time      `json:"harvest_time"`
	Items       map[string]int `json:"items"`
}

// HarvestHistoryDTO 收获历史DTO
type HarvestHistoryDTO struct {
	ID          string    `json:"id"`
	CropType    string    `json:"crop_type"`
	Quantity    int       `json:"quantity"`
	Quality     float64   `json:"quality"`
	Experience  int64     `json:"experience"`
	HarvestTime time.Time `json:"harvest_time"`
}

// PlantingStatisticsDTO 种植统计DTO
type PlantingStatisticsDTO struct {
	PlayerID        string                    `json:"player_id"`
	TotalPlanted    int64                     `json:"total_planted"`
	TotalHarvested  int64                     `json:"total_harvested"`
	TotalExperience int64                     `json:"total_experience"`
	CropTypeStats   map[string]*CropTypeStats `json:"crop_type_stats"`
	AverageQuality  float64                   `json:"average_quality"`
	FavoriteCrop    string                    `json:"favorite_crop"`
	LastPlantTime   time.Time                 `json:"last_plant_time"`
	LastHarvestTime time.Time                 `json:"last_harvest_time"`
}

// CropTypeStats 作物类型统计
type CropTypeStats struct {
	CropType       string  `json:"crop_type"`
	TotalPlanted   int64   `json:"total_planted"`
	TotalHarvested int64   `json:"total_harvested"`
	AverageQuality float64 `json:"average_quality"`
	TotalYield     int64   `json:"total_yield"`
	SuccessRate    float64 `json:"success_rate"`
}
