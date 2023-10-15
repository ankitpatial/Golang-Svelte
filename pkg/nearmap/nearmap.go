package nearmap

const (
	ProviderName = "NearMap"

	precision = uint(3)

	PitchDegreePrimaryRoof = "31225e77-2159-5509-a83c-838c18445dbf"
	TreeOverhangPresent    = "6132706b-622f-5c68-ad85-3b3f7d6b6f5b"

	roofCount     = "51cbb913-1c5c-5253-b333-3789e34839de"
	unclippedArea = "b04d6b59-5307-57d9-a763-dfa64100670a"

	// flat (roof types) area (sq. ft) of primary roof
	flat = "a5d04b2f-7f29-5279-80a1-3403788e1911"
	// flatRatio is flat (roof types) ratio in primary roof
	flatRatio = "9a6dcdc5-ac39-52a7-9f3c-9b473784a655"

	// shingleDominant in primary roof
	shingleDominant = "9ae46172-8a4b-553e-9f36-e325efedfc61"
	// shingle (roof material) area (sq. ft) of primary roof
	shingle = "ea83ca07-b493-5232-a60d-6165d7b8535b"
	// shingleRatio ratio in primary roof
	shingleRatio = "4a5af13b-d1a1-5169-8284-28ffdd1d950c"
	// shingleT is total unclipped area (sq. ft) of shingle
	shingleT = "de592b36-efde-54ac-b532-8a4692ca5c09"

	// tileDominant in primary roof
	tileDominant = "18a30033-ec67-5e98-8c76-918f7bb207c8"
	// tile (roof material) area (sq. ft) of primary roof
	tile = "3be5bbdd-fb8a-56c8-9c05-83224e31323b"
	// tileRatio is tile (roof material) ratio in primary roof
	tileRatio = "ebf8f274-19b4-5a1f-9ebf-f3e3d3f1eee9"
	// tileT is total unclipped area (sq. ft) of tile
	tileT = "e654c744-f3e8-5ec8-a4cd-4157891d9732"

	// metalDominant in primary roof
	metalDominant = "50031e0b-98d1-5524-94b7-bde691eaee76"
	// metal (roof material) area (sq. ft) of primary roof
	metal = "97ff0211-98a5-53f9-a07e-0d0ddc5f8fe5"
	// metalRatio is metal (roof material) ratio in primary roof
	metalRatio = "0491594d-e331-5161-835b-7b3411072c60"
	// metalT is
	metalT = "e59c6f1b-7902-532c-aff2-c9df1461a5bf"

	// gable (roof types) area (sq. ft) of primary roof
	gable = "0e107f3a-1942-5172-91ad-512dc1c19978"
	// gableRatio is gable (roof types) ratio in primary roof
	gableRatio = "58989579-1f7c-50e7-a5a5-33eab090d0c6"
	// gableT total unclipped area (sq. ft) of gable
	gableT = "77fbe573-16fb-5217-9595-b6332b444b36"

	// dutch gable (roof types) area (sq. ft) of primary roof
	dutch = "522cbd32-14ed-5450-b218-19ef92f92edf"
	// dutchRatio is dutch gable (roof types) ratio in primary roof
	dutchRatio = "56a1e713-3675-5bb8-a5f6-d2cd029bd84f"
	// dutchT is total unclipped area (sq. ft) of dutch gable
	dutchT = "6f55d682-fc06-5ad0-aace-333b9edb1d2e"

	// hip (roof types) area (sq. ft) of primary roof
	hip = "2869252c-7972-5228-b620-6c66f24cb964"
	// hipRatio is hip (roof types) ratio in primary roof
	hipRatio = "3c600ebb-0f2e-56e0-8b43-131637a019d5"
	// hipT is total unclipped area (sq. ft) of hip
	hipT = "93849242-f355-5886-8180-322824b0eb9e"

	// turret (roof types) area (sq. ft) of primary roof
	turret = "f4971ff5-e254-5aa7-9c33-16f5939044df"
	// turretRatio is turret (roof types) ratio in primary roof
	turretRatio = "a169a6af-7968-56e6-b47a-66a8308b39ad"
	// turretT is total unclipped area (sq. ft) of turret
	turretT = "ed7effb9-491c-5a38-98a1-695bd7dd768c"

	// treeOHPresent is tree overhang present
	treeOHPresent = "6132706b-622f-5c68-ad85-3b3f7d6b6f5b"
	// treeOHCount is tree overhang count
	treeOHCount = "6196e9dc-db51-5814-858b-b664a1e160a2"
	// treeOH is tree overhang (roof overhang attributes) area (sq. ft) of primary roof
	treeOH = "eeaff46c-9a32-515b-a192-57b026f9cec6"
	// treeOHRatio is tree overhang (roof overhang attributes) ratio in primary roof
	treeOHRatio = "83c0010e-045a-5e35-81bb-6bd0695c1212"
	// treeOHConfidence is  tree overhang (roof overhang attributes) confidence in primary roof
	treeOHConfidence = "8b0cc2ee-c529-5439-b7af-d843b7910b0c"
	// treeOHPConfidence is tree overhang presence confidence
	treeOHPConfidence = "e817a248-b17c-52f2-b6a8-a684f7eab50c"
	// treeOHClipped is total clipped area (sq. ft) of tree overhang
	treeOHClipped = "01b02231-d840-548d-978b-171955619b2b"
	// treeOHUnClipped is total unclipped area (sq. ft) of tree overhang
	treeOHUnClipped = "1d530a3c-efda-54ba-952d-3cf8f1f6f70d"
)

type RoofArea struct {
	SqFt        float64
	Description string
	IsFlat      bool
}

type ResponseAndDetail struct {
	Detail *EstimateDetail        `json:"detail"`
	Raw    map[string]interface{} `json:"raw"`
}
