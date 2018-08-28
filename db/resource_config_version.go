package db

import (
	"database/sql"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/atc"
)

//go:generate counterfeiter . ResourceConfigVersion

type ResourceConfigVersion interface {
	ID() int
	Version() Version
	Metadata() ResourceConfigMetadataFields
	CheckOrder() int
	ResourceConfig() ResourceConfig

	Reload() (bool, error)
	SaveMetadata(ResourceConfigMetadataFields) error
}

type ResourceConfigVersions []ResourceConfigVersion

type ResourceConfigMetadataField struct {
	Name  string
	Value string
}

type ResourceConfigMetadataFields []ResourceConfigMetadataField

func NewResourceConfigMetadataFields(atcm []atc.MetadataField) ResourceConfigMetadataFields {
	metadata := make([]ResourceConfigMetadataField, len(atcm))
	for i, md := range atcm {
		metadata[i] = ResourceConfigMetadataField{
			Name:  md.Name,
			Value: md.Value,
		}
	}

	return metadata
}

type Version map[string]string

type resourceConfigVersion struct {
	id         int
	version    Version
	metadata   ResourceConfigMetadataFields
	checkOrder int

	resourceConfig ResourceConfig

	conn Conn
}

var resourceConfigVersionQuery = psql.Select(`
	id,
	version,
	metadata,
	check_order
`).
	From("resource_config_versions")

func (r *resourceConfigVersion) ID() int                                { return r.id }
func (r *resourceConfigVersion) Version() Version                       { return r.version }
func (r *resourceConfigVersion) Metadata() ResourceConfigMetadataFields { return r.metadata }
func (r *resourceConfigVersion) CheckOrder() int                        { return r.checkOrder }
func (r *resourceConfigVersion) ResourceConfig() ResourceConfig         { return r.resourceConfig }

func (r *resourceConfigVersion) Reload() (bool, error) {
	row := resourceConfigVersionQuery.Where(sq.Eq{"id": r.id}).
		RunWith(r.conn).
		QueryRow()

	err := scanResourceConfigVersion(r, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *resourceConfigVersion) SaveMetadata(metadata ResourceConfigMetadataFields) error {
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	_, err = psql.Update("resource_config_versions").
		Set("metadata", string(metadataJSON)).
		Where(sq.Eq{
			"id": r.id,
		}).
		RunWith(r.conn).
		Exec()

	return err
}

func scanResourceConfigVersion(r *resourceConfigVersion, scan scannable) error {
	var version, metadata sql.NullString

	err := scan.Scan(&r.id, &version, &metadata, &r.checkOrder)
	if err != nil {
		return err
	}

	if version.Valid {
		err = json.Unmarshal([]byte(version.String), &r.version)
		if err != nil {
			return err
		}
	}

	if metadata.Valid {
		err = json.Unmarshal([]byte(metadata.String), &r.metadata)
		if err != nil {
			return err
		}
	}

	return nil
}
