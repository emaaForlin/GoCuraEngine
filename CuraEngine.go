package CuraEngine

import (
    "strconv"
    "os/exec"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "fmt"
    "log"
)

type Config struct {
  Preset []struct{
    Prefix string `json`
    Printer string `json`
    MachineNozzleSize float64 `json`
    MaterialDiameter float64 `json`
    LayerHeight float64 `json`
    Layer0Height float64 `json`
    LineWidth float64 `json`
    WallLineWidth float64 `json`
    WallLineCount int `json`
    OptimizeWallPrintingOrder bool `json`
    SmoothEnable bool `json`
    InfillDensity int `json`
    InfillPattern string `json`
    BedTemp int `json`
    HotendTemp int `json`
    SpeedPrint float64 `json`
    SpeedInfill float64 `json`
    SpeedPrintLayer0 float64 `json`
    SpeedWall float64 `json`
    SpeedTopBottom float64 `json`
    SpeedTravel float64 `json`
    RetractionEnable bool `json`
    RetractionDistance float64 `json`
    RetractionSpeed float64 `json`
    ZHopOnRetract bool `json`
    AvoidOthersOnTravel bool `json`
    AvoidSupportsOnTravel bool `json`
    DistanceAvoid float64 `json`
    CoollingFanEnable bool `json`
    CoolFanSpeed int `json`
    SupportEnable bool `json`
    SupportTreeEnable bool `json`
    AdhesionType string `json`
  }
}

type Params struct {
  Printer string
  MachineNozzleSize string
  MaterialDiameter string
  LayerHeight string
  Layer0Height string
  LineWidth string
  WallLineWidth string
  WallLineCount string
  OptimizeWallPrintingOrder string
  SmoothEnable string
  InfillDensity string
  InfillPattern string
  BedTemp string
  HotendTemp string
  SpeedPrint string
  SpeedInfill string
  SpeedPrintLayer0 string
  SpeedWall string
  SpeedTopBottom string
  SpeedTravel string
  RetractionEnable string
  RetractionDistance string
  RetractionSpeed string
  ZHopOnRetract string
  AvoidOthersOnTravel string
  AvoidSupportsOnTravel string
  DistanceAvoid string
  CoollingFanEnable string
  CoolFanSpeed string
  SupportEnable string
  SupportTreeEnable string
  AdhesionType string
}


func LoadFile(jsonFilePath string) Config {
  jsonData, err := ioutil.ReadFile(jsonFilePath)
  if err != nil {
    log.Fatal(err)
  }
  p := Config{}
  e := json.Unmarshal([]byte(jsonData), &p)
  if e != nil {
      log.Fatal(e)
  }
    log.Println("File successfully loaded.")
    return p
}

func ParseData(config Config, id int) Params {
  var p Params
  p.Printer = config.Preset[id].Printer
  p.MachineNozzleSize = strconv.FormatFloat(config.Preset[id].MachineNozzleSize, 'G', -1, 64)
  p.MaterialDiameter = strconv.FormatFloat(config.Preset[id].MaterialDiameter, 'G', -1, 64)
  p.LayerHeight = strconv.FormatFloat(config.Preset[id].LayerHeight, 'G', -1, 64)
  p.Layer0Height = strconv.FormatFloat(config.Preset[id].Layer0Height, 'G', -1, 64)
  p.LineWidth = strconv.FormatFloat(config.Preset[id].LineWidth, 'G', -1, 64)
  p.WallLineWidth = strconv.FormatFloat(config.Preset[id].WallLineWidth, 'G', -1, 64)
  p.WallLineCount = strconv.Itoa(config.Preset[id].WallLineCount)
  p.OptimizeWallPrintingOrder = strconv.FormatBool(config.Preset[id].OptimizeWallPrintingOrder)
  p.SmoothEnable = strconv.FormatBool(config.Preset[id].SmoothEnable)
  p.InfillDensity = strconv.Itoa(config.Preset[id].InfillDensity)
  p.InfillPattern = config.Preset[id].InfillPattern
  p.BedTemp = strconv.Itoa(config.Preset[id].BedTemp)
  p.HotendTemp = strconv.Itoa(config.Preset[id].HotendTemp)
  p.SpeedPrint = strconv.FormatFloat(config.Preset[id].SpeedPrint, 'G', -1, 64)
  p.SpeedInfill = strconv.FormatFloat(config.Preset[id].SpeedInfill, 'G', -1, 64)
  p.SpeedPrintLayer0 = strconv.FormatFloat(config.Preset[id].SpeedPrintLayer0, 'G', -1, 64)
  p.SpeedWall = strconv.FormatFloat(config.Preset[id].SpeedWall, 'G', -1, 64)
  p.SpeedTopBottom = strconv.FormatFloat(config.Preset[id].SpeedTopBottom, 'G', -1, 64)
  p.SpeedTravel = strconv.FormatFloat(config.Preset[id].SpeedTravel, 'G', -1, 64)
  p.RetractionEnable = strconv.FormatBool(config.Preset[id].RetractionEnable)
  p.RetractionDistance = strconv.FormatFloat(config.Preset[id].RetractionDistance, 'G', -1, 64)
  p.RetractionSpeed = strconv.FormatFloat(config.Preset[id].RetractionSpeed, 'G', -1, 64)
  p.ZHopOnRetract = strconv.FormatBool(config.Preset[id].ZHopOnRetract)
  p.AvoidOthersOnTravel = strconv.FormatBool(config.Preset[id].AvoidOthersOnTravel)
  p.AvoidSupportsOnTravel = strconv.FormatBool(config.Preset[id].AvoidSupportsOnTravel)
  p.DistanceAvoid = strconv.FormatFloat(config.Preset[id].DistanceAvoid, 'G', -1, 64)
  p.CoollingFanEnable = strconv.FormatBool(config.Preset[id].CoollingFanEnable)
  p.CoolFanSpeed = strconv.Itoa(config.Preset[id].CoolFanSpeed)
  p.SupportEnable = strconv.FormatBool(config.Preset[id].SupportEnable)
  p.SupportTreeEnable = strconv.FormatBool(config.Preset[id].SupportTreeEnable)
  p.AdhesionType = config.Preset[id].AdhesionType

  log.Println("Successfully data parsing.")
  return p
}

func Slice(p Params, model string, output string){
  cmd := exec.Command("CuraEngine", "slice", "-p", "-j", p.Printer, "-s", "support_tree_enable="+p.SupportTreeEnable, "-s", "spaghetti_infill_enabled=false", "-s", "machine_nozzle_size="+p.MachineNozzleSize, "-s","material_diameter="+p.MaterialDiameter,"-s", "layer_height="+p.LayerHeight, "-s", "layer_height_0="+p.Layer0Height, "-s", "line_width="+p.LineWidth, "-s", "wall_line_width_x="+p.LineWidth, "-s", "wall_line_count="+p.WallLineCount, "-s", "optimize_wall_printing_order="+p.OptimizeWallPrintingOrder, "-s", "smooth_spiralized_contours="+p.SmoothEnable, "-s", "infill_sparse_density="+p.InfillDensity, "-s", "infill_pattern="+p.InfillPattern, "-s", "material_bed_temperature="+p.BedTemp, "-s", "material_print_temperature="+p.HotendTemp, "-s", "speed_print="+p.SpeedPrint, "-s", "speed_infill="+p.SpeedInfill, "-s", "speed_print_layer_0="+p.SpeedPrintLayer0, "-s", "speed_wall="+p.SpeedWall, "-s", "speed_topbottom="+p.SpeedTopBottom, "-s", "speed_travel="+p.SpeedTravel, "-s",  "retraction_enable="+p.RetractionEnable, "-s", "retraction_speed="+p.RetractionSpeed, "-s", "retraction_hop_enabled="+p.ZHopOnRetract, "-s", "travel_avoid_other_parts="+p.AvoidOthersOnTravel, "-s", "travel_avoid_supports="+p.AvoidSupportsOnTravel, "-s", "cool_fan_enabled="+p.CoollingFanEnable, "-s", "cool_fan_speed_min="+p.CoolFanSpeed, "-s", "support_enable="+p.SupportEnable, "-s", "adhesion_type="+p.AdhesionType, "-l", model,"-o", output)

  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    log.Println(fmt.Sprint(err) + ": " + stderr.String())
    log.Fatal(err)
    return
  }
  log.Println(out.String())
  log.Println("Successfully model slicing!")
}
