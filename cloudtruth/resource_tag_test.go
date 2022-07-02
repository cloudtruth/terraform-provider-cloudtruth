package cloudtruth

/*
const tagDesc = "Just a description of a tag"
const updateTagDesc = "A new description of an tag"
const tagVal = "A useful string"
const updateTagVal = "A new useful string"
const prodTagVal = "A useful string only in production"

func TestAccResourceTagBasic(t *testing.T) {
	createTagName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTagCreateBasic(accTestProject, createTagName, tagDesc, paramVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "name", createTagName),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "description", tagDesc),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "value", tagVal),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "external",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "dynamic",
						strconv.FormatBool(false)),
				),
			},
			{
				Config: testAccResourceTagUpdateBasic(accTestProject, createTagName, updateParamDesc, updateParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "name", createTagName),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "description", updateTagDesc),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "value", updateTagVal),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "external",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "dynamic",
						strconv.FormatBool(false)),
				),
			},
		},
	})
}
*/
