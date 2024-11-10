package models

import "time"

type CraftReq struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type UnEquipReq struct {
	Slot string `json:"slot"`
}

type EquipReq struct {
	Code string `json:"code"`
	Slot string `json:"slot"`
}

type MoveReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type BuyReq struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type SellReq struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type Equip struct {
	Data struct {
		Cooldown struct {
			TotalSeconds     int       `json:"total_seconds"`
			RemainingSeconds int       `json:"remaining_seconds"`
			StartedAt        time.Time `json:"started_at"`
			Expiration       time.Time `json:"expiration"`
			Reason           string    `json:"reason"`
		} `json:"cooldown"`
		Slot string `json:"slot"`
		Item struct {
			Name        string `json:"name"`
			Code        string `json:"code"`
			Level       int    `json:"level"`
			Type        string `json:"type"`
			Subtype     string `json:"subtype"`
			Description string `json:"description"`
			Effects     []struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
			} `json:"effects"`
			Craft struct {
				Skill string `json:"skill"`
				Level int    `json:"level"`
				Items []struct {
					Code     string `json:"code"`
					Quantity int    `json:"quantity"`
				} `json:"items"`
				Quantity int `json:"quantity"`
			} `json:"craft"`
		} `json:"item"`
		Character struct {
			Name                    string    `json:"name"`
			Skin                    string    `json:"skin"`
			Level                   int       `json:"level"`
			Xp                      int       `json:"xp"`
			MaxXp                   int       `json:"max_xp"`
			TotalXp                 int       `json:"total_xp"`
			Gold                    int       `json:"gold"`
			Speed                   int       `json:"speed"`
			MiningLevel             int       `json:"mining_level"`
			MiningXp                int       `json:"mining_xp"`
			MiningMaxXp             int       `json:"mining_max_xp"`
			WoodcuttingLevel        int       `json:"woodcutting_level"`
			WoodcuttingXp           int       `json:"woodcutting_xp"`
			WoodcuttingMaxXp        int       `json:"woodcutting_max_xp"`
			FishingLevel            int       `json:"fishing_level"`
			FishingXp               int       `json:"fishing_xp"`
			FishingMaxXp            int       `json:"fishing_max_xp"`
			WeaponcraftingLevel     int       `json:"weaponcrafting_level"`
			WeaponcraftingXp        int       `json:"weaponcrafting_xp"`
			WeaponcraftingMaxXp     int       `json:"weaponcrafting_max_xp"`
			GearcraftingLevel       int       `json:"gearcrafting_level"`
			GearcraftingXp          int       `json:"gearcrafting_xp"`
			GearcraftingMaxXp       int       `json:"gearcrafting_max_xp"`
			JewelrycraftingLevel    int       `json:"jewelrycrafting_level"`
			JewelrycraftingXp       int       `json:"jewelrycrafting_xp"`
			JewelrycraftingMaxXp    int       `json:"jewelrycrafting_max_xp"`
			CookingLevel            int       `json:"cooking_level"`
			CookingXp               int       `json:"cooking_xp"`
			CookingMaxXp            int       `json:"cooking_max_xp"`
			Hp                      int       `json:"hp"`
			Haste                   int       `json:"haste"`
			CriticalStrike          int       `json:"critical_strike"`
			Stamina                 int       `json:"stamina"`
			AttackFire              int       `json:"attack_fire"`
			AttackEarth             int       `json:"attack_earth"`
			AttackWater             int       `json:"attack_water"`
			AttackAir               int       `json:"attack_air"`
			DmgFire                 int       `json:"dmg_fire"`
			DmgEarth                int       `json:"dmg_earth"`
			DmgWater                int       `json:"dmg_water"`
			DmgAir                  int       `json:"dmg_air"`
			ResFire                 int       `json:"res_fire"`
			ResEarth                int       `json:"res_earth"`
			ResWater                int       `json:"res_water"`
			ResAir                  int       `json:"res_air"`
			X                       int       `json:"x"`
			Y                       int       `json:"y"`
			Cooldown                int       `json:"cooldown"`
			CooldownExpiration      time.Time `json:"cooldown_expiration"`
			WeaponSlot              string    `json:"weapon_slot"`
			ShieldSlot              string    `json:"shield_slot"`
			HelmetSlot              string    `json:"helmet_slot"`
			BodyArmorSlot           string    `json:"body_armor_slot"`
			LegArmorSlot            string    `json:"leg_armor_slot"`
			BootsSlot               string    `json:"boots_slot"`
			Ring1Slot               string    `json:"ring1_slot"`
			Ring2Slot               string    `json:"ring2_slot"`
			AmuletSlot              string    `json:"amulet_slot"`
			Artifact1Slot           string    `json:"artifact1_slot"`
			Artifact2Slot           string    `json:"artifact2_slot"`
			Artifact3Slot           string    `json:"artifact3_slot"`
			Consumable1Slot         string    `json:"consumable1_slot"`
			Consumable1SlotQuantity int       `json:"consumable1_slot_quantity"`
			Consumable2Slot         string    `json:"consumable2_slot"`
			Consumable2SlotQuantity int       `json:"consumable2_slot_quantity"`
			Task                    string    `json:"task"`
			TaskType                string    `json:"task_type"`
			TaskProgress            int       `json:"task_progress"`
			TaskTotal               int       `json:"task_total"`
			InventoryMaxItems       int       `json:"inventory_max_items"`
			Inventory               []struct {
				Slot     int    `json:"slot"`
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"inventory"`
		} `json:"character"`
	} `json:"data"`
}

type Transaction struct {
	Data struct {
		Cooldown struct {
			TotalSeconds     int       `json:"total_seconds"`
			RemainingSeconds int       `json:"remaining_seconds"`
			StartedAt        time.Time `json:"started_at"`
			Expiration       time.Time `json:"expiration"`
			Reason           string    `json:"reason"`
		} `json:"cooldown"`
		Transaction struct {
			Code       string `json:"code"`
			Quantity   int    `json:"quantity"`
			Price      int    `json:"price"`
			TotalPrice int    `json:"total_price"`
		} `json:"transaction"`
		Character struct {
			Name                    string    `json:"name"`
			Skin                    string    `json:"skin"`
			Level                   int       `json:"level"`
			Xp                      int       `json:"xp"`
			MaxXp                   int       `json:"max_xp"`
			TotalXp                 int       `json:"total_xp"`
			Gold                    int       `json:"gold"`
			Speed                   int       `json:"speed"`
			MiningLevel             int       `json:"mining_level"`
			MiningXp                int       `json:"mining_xp"`
			MiningMaxXp             int       `json:"mining_max_xp"`
			WoodcuttingLevel        int       `json:"woodcutting_level"`
			WoodcuttingXp           int       `json:"woodcutting_xp"`
			WoodcuttingMaxXp        int       `json:"woodcutting_max_xp"`
			FishingLevel            int       `json:"fishing_level"`
			FishingXp               int       `json:"fishing_xp"`
			FishingMaxXp            int       `json:"fishing_max_xp"`
			WeaponcraftingLevel     int       `json:"weaponcrafting_level"`
			WeaponcraftingXp        int       `json:"weaponcrafting_xp"`
			WeaponcraftingMaxXp     int       `json:"weaponcrafting_max_xp"`
			GearcraftingLevel       int       `json:"gearcrafting_level"`
			GearcraftingXp          int       `json:"gearcrafting_xp"`
			GearcraftingMaxXp       int       `json:"gearcrafting_max_xp"`
			JewelrycraftingLevel    int       `json:"jewelrycrafting_level"`
			JewelrycraftingXp       int       `json:"jewelrycrafting_xp"`
			JewelrycraftingMaxXp    int       `json:"jewelrycrafting_max_xp"`
			CookingLevel            int       `json:"cooking_level"`
			CookingXp               int       `json:"cooking_xp"`
			CookingMaxXp            int       `json:"cooking_max_xp"`
			Hp                      int       `json:"hp"`
			Haste                   int       `json:"haste"`
			CriticalStrike          int       `json:"critical_strike"`
			Stamina                 int       `json:"stamina"`
			AttackFire              int       `json:"attack_fire"`
			AttackEarth             int       `json:"attack_earth"`
			AttackWater             int       `json:"attack_water"`
			AttackAir               int       `json:"attack_air"`
			DmgFire                 int       `json:"dmg_fire"`
			DmgEarth                int       `json:"dmg_earth"`
			DmgWater                int       `json:"dmg_water"`
			DmgAir                  int       `json:"dmg_air"`
			ResFire                 int       `json:"res_fire"`
			ResEarth                int       `json:"res_earth"`
			ResWater                int       `json:"res_water"`
			ResAir                  int       `json:"res_air"`
			X                       int       `json:"x"`
			Y                       int       `json:"y"`
			Cooldown                int       `json:"cooldown"`
			CooldownExpiration      time.Time `json:"cooldown_expiration"`
			WeaponSlot              string    `json:"weapon_slot"`
			ShieldSlot              string    `json:"shield_slot"`
			HelmetSlot              string    `json:"helmet_slot"`
			BodyArmorSlot           string    `json:"body_armor_slot"`
			LegArmorSlot            string    `json:"leg_armor_slot"`
			BootsSlot               string    `json:"boots_slot"`
			Ring1Slot               string    `json:"ring1_slot"`
			Ring2Slot               string    `json:"ring2_slot"`
			AmuletSlot              string    `json:"amulet_slot"`
			Artifact1Slot           string    `json:"artifact1_slot"`
			Artifact2Slot           string    `json:"artifact2_slot"`
			Artifact3Slot           string    `json:"artifact3_slot"`
			Consumable1Slot         string    `json:"consumable1_slot"`
			Consumable1SlotQuantity int       `json:"consumable1_slot_quantity"`
			Consumable2Slot         string    `json:"consumable2_slot"`
			Consumable2SlotQuantity int       `json:"consumable2_slot_quantity"`
			Task                    string    `json:"task"`
			TaskType                string    `json:"task_type"`
			TaskProgress            int       `json:"task_progress"`
			TaskTotal               int       `json:"task_total"`
			InventoryMaxItems       int       `json:"inventory_max_items"`
			Inventory               []struct {
				Slot     int    `json:"slot"`
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"inventory"`
		} `json:"character"`
	} `json:"data"`
}

type Fight struct {
	Data struct {
		Cooldown struct {
			TotalSeconds     int       `json:"total_seconds"`
			RemainingSeconds int       `json:"remaining_seconds"`
			StartedAt        time.Time `json:"started_at"`
			Expiration       time.Time `json:"expiration"`
			Reason           string    `json:"reason"`
		} `json:"cooldown"`
		Fight struct {
			Xp    int `json:"xp"`
			Gold  int `json:"gold"`
			Drops []struct {
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"drops"`
			Turns              int `json:"turns"`
			MonsterBlockedHits struct {
				Fire  int `json:"fire"`
				Earth int `json:"earth"`
				Water int `json:"water"`
				Air   int `json:"air"`
				Total int `json:"total"`
			} `json:"monster_blocked_hits"`
			PlayerBlockedHits struct {
				Fire  int `json:"fire"`
				Earth int `json:"earth"`
				Water int `json:"water"`
				Air   int `json:"air"`
				Total int `json:"total"`
			} `json:"player_blocked_hits"`
			Logs   []string `json:"logs"`
			Result string   `json:"result"`
		} `json:"fight"`
		Character struct {
			Name                 string    `json:"name"`
			Account              string    `json:"account"`
			Skin                 string    `json:"skin"`
			Level                int       `json:"level"`
			Xp                   int       `json:"xp"`
			MaxXp                int       `json:"max_xp"`
			Gold                 int       `json:"gold"`
			Speed                int       `json:"speed"`
			MiningLevel          int       `json:"mining_level"`
			MiningXp             int       `json:"mining_xp"`
			MiningMaxXp          int       `json:"mining_max_xp"`
			WoodcuttingLevel     int       `json:"woodcutting_level"`
			WoodcuttingXp        int       `json:"woodcutting_xp"`
			WoodcuttingMaxXp     int       `json:"woodcutting_max_xp"`
			FishingLevel         int       `json:"fishing_level"`
			FishingXp            int       `json:"fishing_xp"`
			FishingMaxXp         int       `json:"fishing_max_xp"`
			WeaponcraftingLevel  int       `json:"weaponcrafting_level"`
			WeaponcraftingXp     int       `json:"weaponcrafting_xp"`
			WeaponcraftingMaxXp  int       `json:"weaponcrafting_max_xp"`
			GearcraftingLevel    int       `json:"gearcrafting_level"`
			GearcraftingXp       int       `json:"gearcrafting_xp"`
			GearcraftingMaxXp    int       `json:"gearcrafting_max_xp"`
			JewelrycraftingLevel int       `json:"jewelrycrafting_level"`
			JewelrycraftingXp    int       `json:"jewelrycrafting_xp"`
			JewelrycraftingMaxXp int       `json:"jewelrycrafting_max_xp"`
			CookingLevel         int       `json:"cooking_level"`
			CookingXp            int       `json:"cooking_xp"`
			CookingMaxXp         int       `json:"cooking_max_xp"`
			AlchemyLevel         int       `json:"alchemy_level"`
			AlchemyXp            int       `json:"alchemy_xp"`
			AlchemyMaxXp         int       `json:"alchemy_max_xp"`
			Hp                   int       `json:"hp"`
			MaxHp                int       `json:"max_hp"`
			Haste                int       `json:"haste"`
			CriticalStrike       int       `json:"critical_strike"`
			Stamina              int       `json:"stamina"`
			AttackFire           int       `json:"attack_fire"`
			AttackEarth          int       `json:"attack_earth"`
			AttackWater          int       `json:"attack_water"`
			AttackAir            int       `json:"attack_air"`
			DmgFire              int       `json:"dmg_fire"`
			DmgEarth             int       `json:"dmg_earth"`
			DmgWater             int       `json:"dmg_water"`
			DmgAir               int       `json:"dmg_air"`
			ResFire              int       `json:"res_fire"`
			ResEarth             int       `json:"res_earth"`
			ResWater             int       `json:"res_water"`
			ResAir               int       `json:"res_air"`
			X                    int       `json:"x"`
			Y                    int       `json:"y"`
			Cooldown             int       `json:"cooldown"`
			CooldownExpiration   time.Time `json:"cooldown_expiration"`
			WeaponSlot           string    `json:"weapon_slot"`
			ShieldSlot           string    `json:"shield_slot"`
			HelmetSlot           string    `json:"helmet_slot"`
			BodyArmorSlot        string    `json:"body_armor_slot"`
			LegArmorSlot         string    `json:"leg_armor_slot"`
			BootsSlot            string    `json:"boots_slot"`
			Ring1Slot            string    `json:"ring1_slot"`
			Ring2Slot            string    `json:"ring2_slot"`
			AmuletSlot           string    `json:"amulet_slot"`
			Artifact1Slot        string    `json:"artifact1_slot"`
			Artifact2Slot        string    `json:"artifact2_slot"`
			Artifact3Slot        string    `json:"artifact3_slot"`
			Utility1Slot         string    `json:"utility1_slot"`
			Utility1SlotQuantity int       `json:"utility1_slot_quantity"`
			Utility2Slot         string    `json:"utility2_slot"`
			Utility2SlotQuantity int       `json:"utility2_slot_quantity"`
			Task                 string    `json:"task"`
			TaskType             string    `json:"task_type"`
			TaskProgress         int       `json:"task_progress"`
			TaskTotal            int       `json:"task_total"`
			InventoryMaxItems    int       `json:"inventory_max_items"`
			Inventory            []struct {
				Slot     int    `json:"slot"`
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"inventory"`
		} `json:"character"`
	} `json:"data"`
}

type Action struct {
	Data struct {
		Cooldown struct {
			TotalSeconds     int       `json:"total_seconds"`
			RemainingSeconds int       `json:"remaining_seconds"`
			StartedAt        time.Time `json:"started_at"`
			Expiration       time.Time `json:"expiration"`
			Reason           string    `json:"reason"`
		} `json:"cooldown"`
		Destination struct {
			Name    string `json:"name"`
			Skin    string `json:"skin"`
			X       int    `json:"x"`
			Y       int    `json:"y"`
			Content struct {
				Type string `json:"type"`
				Code string `json:"code"`
			} `json:"content"`
		} `json:"destination"`
		Character struct {
			Name                 string    `json:"name"`
			Account              string    `json:"account"`
			Skin                 string    `json:"skin"`
			Level                int       `json:"level"`
			Xp                   int       `json:"xp"`
			MaxXp                int       `json:"max_xp"`
			Gold                 int       `json:"gold"`
			Speed                int       `json:"speed"`
			MiningLevel          int       `json:"mining_level"`
			MiningXp             int       `json:"mining_xp"`
			MiningMaxXp          int       `json:"mining_max_xp"`
			WoodcuttingLevel     int       `json:"woodcutting_level"`
			WoodcuttingXp        int       `json:"woodcutting_xp"`
			WoodcuttingMaxXp     int       `json:"woodcutting_max_xp"`
			FishingLevel         int       `json:"fishing_level"`
			FishingXp            int       `json:"fishing_xp"`
			FishingMaxXp         int       `json:"fishing_max_xp"`
			WeaponcraftingLevel  int       `json:"weaponcrafting_level"`
			WeaponcraftingXp     int       `json:"weaponcrafting_xp"`
			WeaponcraftingMaxXp  int       `json:"weaponcrafting_max_xp"`
			GearcraftingLevel    int       `json:"gearcrafting_level"`
			GearcraftingXp       int       `json:"gearcrafting_xp"`
			GearcraftingMaxXp    int       `json:"gearcrafting_max_xp"`
			JewelrycraftingLevel int       `json:"jewelrycrafting_level"`
			JewelrycraftingXp    int       `json:"jewelrycrafting_xp"`
			JewelrycraftingMaxXp int       `json:"jewelrycrafting_max_xp"`
			CookingLevel         int       `json:"cooking_level"`
			CookingXp            int       `json:"cooking_xp"`
			CookingMaxXp         int       `json:"cooking_max_xp"`
			AlchemyLevel         int       `json:"alchemy_level"`
			AlchemyXp            int       `json:"alchemy_xp"`
			AlchemyMaxXp         int       `json:"alchemy_max_xp"`
			Hp                   int       `json:"hp"`
			MaxHp                int       `json:"max_hp"`
			Haste                int       `json:"haste"`
			CriticalStrike       int       `json:"critical_strike"`
			Stamina              int       `json:"stamina"`
			AttackFire           int       `json:"attack_fire"`
			AttackEarth          int       `json:"attack_earth"`
			AttackWater          int       `json:"attack_water"`
			AttackAir            int       `json:"attack_air"`
			DmgFire              int       `json:"dmg_fire"`
			DmgEarth             int       `json:"dmg_earth"`
			DmgWater             int       `json:"dmg_water"`
			DmgAir               int       `json:"dmg_air"`
			ResFire              int       `json:"res_fire"`
			ResEarth             int       `json:"res_earth"`
			ResWater             int       `json:"res_water"`
			ResAir               int       `json:"res_air"`
			X                    int       `json:"x"`
			Y                    int       `json:"y"`
			Cooldown             int       `json:"cooldown"`
			CooldownExpiration   time.Time `json:"cooldown_expiration"`
			WeaponSlot           string    `json:"weapon_slot"`
			ShieldSlot           string    `json:"shield_slot"`
			HelmetSlot           string    `json:"helmet_slot"`
			BodyArmorSlot        string    `json:"body_armor_slot"`
			LegArmorSlot         string    `json:"leg_armor_slot"`
			BootsSlot            string    `json:"boots_slot"`
			Ring1Slot            string    `json:"ring1_slot"`
			Ring2Slot            string    `json:"ring2_slot"`
			AmuletSlot           string    `json:"amulet_slot"`
			Artifact1Slot        string    `json:"artifact1_slot"`
			Artifact2Slot        string    `json:"artifact2_slot"`
			Artifact3Slot        string    `json:"artifact3_slot"`
			Utility1Slot         string    `json:"utility1_slot"`
			Utility1SlotQuantity int       `json:"utility1_slot_quantity"`
			Utility2Slot         string    `json:"utility2_slot"`
			Utility2SlotQuantity int       `json:"utility2_slot_quantity"`
			Task                 string    `json:"task"`
			TaskType             string    `json:"task_type"`
			TaskProgress         int       `json:"task_progress"`
			TaskTotal            int       `json:"task_total"`
			InventoryMaxItems    int       `json:"inventory_max_items"`
			Inventory            []struct {
				Slot     int    `json:"slot"`
				Code     string `json:"code"`
				Quantity int    `json:"quantity"`
			} `json:"inventory"`
		} `json:"character"`
	} `json:"data"`
}
